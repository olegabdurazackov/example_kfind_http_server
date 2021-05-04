package main

import (
	"database/sql"
//	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
//	"os/exec"
  "strings"
  spl "libs_m/word"
//  cl "libs_m/color"
  "net/http"
  "os"
)
var version = "release 0.00.0"
type Doc struct {
    doc_id  int64
    numb    string
    name    string
    abstract     string
    dstart     string
    dend    string
    country     string
    filepath string
    filename string
    doc_mem string
    key string
}

var tmp1/*, tmp2, tmp3*/ string
var start=false

func find_page(w http.ResponseWriter, r *http.Request ){

    var head=`<!doctype html><html lang="ru">
    <head>
    <title>Поиск ГОСТа или СНиПа</title>
    <meta name="generator" content="Bluefish 2.2.6" >
    <meta name="author" content="o" >
    <meta name="date" content="2021-02-15T09:41:02+0500" >
    <meta name="copyright" content="">
    <meta name="keywords" content="">
    <meta name="description" content="">
    <meta name="ROBOTS" content="NOINDEX, NOFOLLOW">
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">
    <meta http-equiv="content-type" content="application/xhtml+xml; charset=UTF-8">
    <meta http-equiv="content-style-type" content="text/css">
    <meta http-equiv="expires" content="0">
    <style type="text/css">
    h1,h2,h3 {color: maroon; font-family: sans-serif; text-shadow: green 1px 2px;}
      h4 { padding-left: 2em; padding-right: 2em; font-family: serif; line-height: 1.5;}
      body {background: #DFF;font-family: serif;  }
      p {color: green; text-indent: 4em;padding-left: 4em; padding-right: 4em; line-height: 1.5;}
      form {padding-left: 2em; padding-right: 2em;}
    </style>
    </head>
    <body>

    `
    var form=`<form action="/snips" method="post" >
    <hgroup><h1>Поиск информации</h1> <h2>в СНиПах и ГОСТах</h2></hgroup>
    <fieldset><legend>основная информация</legend>
    <div>
    <label for="number">номер (шифр)</label>
    <input id="number" type="text" name="number" placeholder="Например: ГОСТ Р 123-45">
    </div>
    <div>
    <label for="title">название </label>
    <input id="title" type="text" name="title" placeholder="Например: электрические сети">
    </div>
    <div>
    <label for="jahr">год ввода в действие </label>
    <input id="jahr" type="text" name="jahr" pattern="[1-2][0-9]{3}"  placeholder="Например: 1990">
    </div></fieldset>
    <fieldset><legend>дополнительная информация</legend>
    <div>
    <label for="abs">аннотация </label>
    <input id="abs" type="text" name="abs" placeholder="Например: введен впервые">
    </div>
    <div>
    <label for="key">ключевое слово </label>
    <input id="key" type="text" name="key" placeholder="Например: любые слова">
    </div></fieldset>
    <fieldset><legend>а так же</legend>
    <div>
    <label for="do">ввод до </label>
    <input id="do" type="text" name="do"  pattern="[1-2][0-9]{3}" value="2999">
    </div>
    <div>
    <label for="posle">ввод позже </label>
    <input id="posle" type="text" name="posle"  pattern="[1-2][0-9]{3}" value="1700">
    </div>
    <div>
    <label for="end">год вывода </label>
    <input id="end" type="text" name="end" placeholder="прекращения действия">
    </div>
    <div>
    <label for="country">страна </label>
    <input id="country" type="text" name="country" pattern="[a-z][a-z]"  list="udssr">
    <datalist id="udssr">
    <select><option value="ru"></option>
    <option value="su"></option></select></datalist>
    </div></fieldset>
    <button name="short" value="0" type="submit">Полный вывод</button>
    <button name="short" value="1" type="submit">Краткий вывод</button>
    </form>
    `
    var foot=`</body></html>`
    var an_err=`<p class="error">%s</p>`
    var  flnumb,flname,fldo,fljahr ,flend ,flposle , flabstract, flkeyword , flcountry , flshort string 
  err:=r.ParseForm()
  fmt.Fprint(w,head,form)
  if err!=nil{
      fmt.Fprintf(w,an_err,err)
  }
  if !start {
      fmt.Fprintf(w,foot)
  }else{
      flnumb   =r.Form["number"][0]
      flname   =r.Form["title"][0] 
      fldo     =r.Form["do"][0]
      fljahr   =r.Form["jahr"][0]
      flend    =r.Form["end"][0]
      flposle  =r.Form["posle"][0]
      flabstract=r.Form["abs"][0]
      flkeyword= r.Form["key"] [0]
      flcountry=r.Form["country"][0]
      flshort  =r.Form["short"][0]
  }
  fmt.Println( start,flnumb,flname,fldo,fljahr ,flend ,flposle , flabstract, flkeyword , flcountry , flshort) 
  start=true

	db, err := sql.Open("sqlite3", "db/snips.dbl")
	if err != nil {
		log.Fatal(err)
	}
  println("basa open")
// приведение к sql
  if flnumb!="%"  {flnumb  ="%"+flnumb  +"%"}
  if flname!="%"   {flname   ="%"+flname   +"%"}
  if flcountry!="%"{flcountry="%"+flcountry   +"%"}
  if fljahr!="%"   {fljahr   ="%"+fljahr    +"%"}
  if flend!="%"    {flend    ="%"+flend    +"%"}
  if flabstract!="%"{flabstract="%"+flabstract+"%"}
  if flkeyword!="%" {flkeyword ="%"+flkeyword +"%"}
  flabstract=strings.Replace(flabstract," ","%",-1)

	//    начало отбора
  var slct string
  if flkeyword=="%"{
	     } else {
      str:=flkeyword
      z:=make([]string,0)
      spl.Split_to_word(&str,&z)
      n:=len(z)
      s:="%"
      for i:=0;i<n;i++{
          s=s+z[i]+"%"
      }
      println(s)
      flkeyword=s
  }
  slct = "select * from doc where ( numb like $1 and name like $2 and abstract like $3 and dend like $4 and country like $5 and dstart<$6 and dstart like $7 and dstart>$8 and key like $9 ) order by numb;"

	rows, err := db.Query(slct, flnumb, flname, flabstract, flend, flcountry, fldo, fljahr, flposle, flkeyword)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bks := make([]*Doc, 0)
	for rows.Next() {
		bk := new(Doc)
		err := rows.Scan(&bk.doc_id, &bk.numb, &bk.name, &bk.abstract, &bk.dstart, &bk.dend, &bk.country,  &bk.filepath, &bk.filename, &bk.doc_mem,&tmp1)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// печать

	if flshort == "0" {
		for _, bk := range bks {
        fmt.Fprintf(w,`<p style="color:red;" >%s</p>`,"* * *" )
        fmt.Fprintf(w,`<fieldset><legend><h4>%s,</h4></legend>`,bk.numb)
        fmt.Fprintf(w,`<a href="%s"><h4>%s</h4></a>`,bk.filename,bk.name)
        fmt.Fprintf(w,"<p> %s ( %s -%s )",bk.country,bk.dstart,bk.dend)
        fmt.Fprintf(w,` : <span style="font-size: smaller;">%d *%s</span>`,bk.doc_id,bk.doc_mem)
        fmt.Fprintf(w,`<p> %s</p>`,bk.abstract)
        fmt.Fprintf(w,"%s","</fieldset>")
        //fmt.Fprintf(w,"<p> %s%s",bk.filepath,bk.filename)
		}
	} else {
		for _, bk := range bks {
        fmt.Fprintf(w,`<p style="color:red;" >%s</p>`,"* * *" )
        fmt.Fprintf(w,`<a href="%s"><h4>%s,</h4>`,bk.filename,bk.numb)
        fmt.Fprintf(w,`<h4>%s</h4></a>`,bk.name)
        fmt.Fprintf(w,"<p> %s ( %s -%s )",bk.country,bk.dstart,bk.dend)
        fmt.Fprintf(w,` : <span style="font-size: smaller;">%d</span>`,bk.doc_id)
	  }
fmt.Println()

  fmt.Fprint(w,foot)
  }
}
/*
func home_page(w http.ResponseWriter, r *http.Request ){
    find_page(w,r)
}
*/
func main() {
    port:=":8080"
    if len(os.Args)==2{
        port=":"+os.Args[1]
    }
    println(port)
    fs:=http.FileServer(http.Dir("static"))
    http.Handle("/",fs)
    http.HandleFunc("/snips",find_page)
 //   http.HandleFunc("/",home_page)
    errr:=http.ListenAndServe(port,nil)
    if errr !=nil{
        log.Fatal("ListenAndServei: ",errr)
    }
}
