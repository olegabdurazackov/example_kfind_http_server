package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os/exec"
  "strings"
  spl "libs_m/word"
  cl "libs_m/color"
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

func main() {
  flnumb := flag.String("n", "%",   "поиск номера ")
  flname := flag.String("t", "%",   "поиск названия")
	fldo := flag.String("do", "9999", "поиск по году <")
	fljahr := flag.String("j", "%", "поиск по году =")
	flend := flag.String("e", "%", "поиск по году вывода")
	flposle := flag.String("posle", "0", "поиск по году >")
	flabstract := flag.String("a", "%", "поиск по аннотации")
	flkeyword := flag.String("k", "%", "поиск по ключевому слову")
	flcountry := flag.String("c", "%", "поиск по стране")
	flopen := flag.Int64("open", 0, "открыть по id")
  flshort := flag.String("short", "0", "short [0-1]")
	flag.Parse()
	println("поиск по номеру      -author :", *flnumb)
	println("поиск по названию         -t :", *flname)
	println("поиск по году <          -do :", *fldo)
	println("поиск по году =            -j:", *fljahr)
	println("поиск по году >       -posle :", *flposle)
	println("поиск по году вывода       -e:", *flend)
	println("поиск по стране           -c :", *flcountry)
	println("поиск по аннотации        -a :", *flabstract)
	println("поиск по ключевому слову  -k :", *flkeyword)
	println("открыть по id          -open :", *flopen)
	println("краткий вывод         -short :", *flshort)
	println()

	db, err := sql.Open("sqlite3", "db/snips.dbl")
	if err != nil {
		log.Fatal(err)
	}
// приведение к sql
  if *flnumb!="%"  {*flnumb  ="%"+*flnumb  +"%"}
  if *flname!="%"   {*flname   ="%"+*flname   +"%"}
  if *flcountry!="%"   {*flcountry   ="%"+*flcountry   +"%"}
  if *fljahr!="%"    {*fljahr    ="%"+*fljahr    +"%"}
  if *flend!="%"    {*flend    ="%"+*flend    +"%"}
  if *flabstract!="%"{*flabstract="%"+*flabstract+"%"}
  if *flkeyword!="%" {*flkeyword ="%"+*flkeyword +"%"}
  //*flkeyword=strings.Replace(*flkeyword," ","%",-1)
  *flabstract=strings.Replace(*flabstract," ","%",-1)
//  *floglav=strings.Replace(*floglav," ","%",-1)
//  *flindex=strings.Replace(*flindex," ","%",-1)
  *flname=strings.Replace(*flname," ","%",-1)

	//    начало отбора
  var slct string
  if *flkeyword=="%"{
	     } else {
      str:=*flkeyword
      z:=make([]string,0)
      spl.Split_to_word(&str,&z)
      n:=len(z)
      s:="%"
      for i:=0;i<n;i++{
          s=s+z[i]+"%"
      }
      println(s)
      *flkeyword=s
//      *flabstract=*flkeyword
//    *floglav=*flkeyword
//    *flindex=*flkeyword
//    slct = "select * from abuch where ( numb like $1 and name like $2 and abstract like $3 and dend like $4 and country like $5 and jahr<$6 and jahr like $7 and jahr>$8 and  ( abstract like $9 or key like $10 ));"
  }
  slct = "select * from doc where ( numb like $1 and name like $2 and abstract like $3 and dend like $4 and country like $5 and dstart<$6 and dstart like $7 and dstart>$8 and key like $9 ) order by numb;"

	rows, err := db.Query(slct, *flnumb, *flname, *flabstract, *flend, *flcountry, *fldo, *fljahr, *flposle, *flkeyword)
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

	if *flshort == "0" {
		for _, bk := range bks {
        cl.Set_color(cl.F_blue)
        fmt.Printf("\n\t\t\t\t\t%s","* * *" )
        cl.Set_color(cl.F_magenta)
        fmt.Printf("\n\t%s\n",bk.numb)
        cl.Set_color(cl.B_red)
        cl.Set_color(cl.F_yellow)
        cl.Set_color(cl.Bold)
        fmt.Printf("%d",bk.doc_id)
        cl.Reset_color()
        cl.Set_color(cl.F_cyan)
        fmt.Printf("\t%s",bk.name)
        cl.Set_color(cl.F_red)
        fmt.Printf(" - %s ( %s -%s )",bk.country,bk.dstart,bk.dend)
        cl.Set_color(cl.F_blue)
        fmt.Printf("\n\t%s",bk.doc_mem)
        cl.Set_color(cl.F_green)
        fmt.Printf("\n%s",bk.abstract)
        cl.Set_color(cl.F_blue)
        fmt.Printf("\t%s%s\n",bk.filepath,bk.filename)
        cl.Reset_color()

//			fmt.Printf("id=%d\t%s\n\t%s\n%s\t%s\t%s\t%s  %s\n\tИздательство: %s\n%s\n%s %s\n\n\n", bk.buch_id, bk.author, bk.title, bk.size, bk.jahr, bk.izdan, bk.isbn, bk.memo, bk.izdatel, bk.atext, bk.filepath, bk.filename)
			if *flopen == bk.doc_id {
				tip := bk.filename[len(bk.filename)-4 : len(bk.filename)]
				viewer := "firefox"
				switch tip {
				case ".pdf":
					viewer = "qpdfview"
				case "djvu":
					viewer = "qpdfview"
				case ".fb2":
					viewer = "cr3"
				case "epub":
					viewer = "cr3"
				case ".zip":
					viewer = "cr3"
				case ".chm":
					viewer = "xchm"
				}
				cmd := exec.Command(viewer, (bk.filepath + bk.filename))
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	} else {
		for _, bk := range bks {
        cl.Set_color(cl.F_blue)
        fmt.Printf("\n\t\t\t\t\t%s","* * *" )
        cl.Set_color(cl.F_magenta)
        fmt.Printf("\n\t%s\n",bk.numb)
        cl.Set_color(cl.B_red)
        cl.Set_color(cl.F_yellow)
        cl.Set_color(cl.Bold)
        fmt.Printf("%d",bk.doc_id)
        cl.Reset_color()
        cl.Set_color(cl.F_cyan)
        fmt.Printf("\t%s",bk.name)
        cl.Set_color(cl.F_red)
        fmt.Printf(" - %s ( %s -%s )",bk.country,bk.dstart,bk.dend)
        //fmt.Printf(" - %s - %s, %s ",bk.izdatel,bk.jahr,bk.size)
/*        cl.Set_color(cl.F_blue)
        fmt.Printf("\n\t%s",bk.isbn)
        cl.Set_color(cl.F_green)
        fmt.Printf("\n%s",bk.atext)
        cl.Set_color(cl.F_blue)
        fmt.Printf("\t%s/%s\n",bk.filepath,bk.filename)
*/
        cl.Reset_color()

	//		fmt.Printf("id = %d\t%s\n%s\n%s\t%s\t%s\t%s  %s\nИздательство: %s\n%s %s\n\n", bk.buch_id, bk.author, bk.title, bk.size, bk.jahr, bk.izdan, bk.isbn, bk.memo, bk.izdatel, bk.filepath, bk.filename)
			if *flopen == bk.doc_id {
				tip := bk.filename[len(bk.filename)-4 : len(bk.filename)]
				viewer := "firefox"
				//          println(tip)
				switch tip {
				case ".pdf":
					viewer = "qpdfview"
				case "djvu":
					viewer = "qpdfview"
				case ".fb2":
					viewer = "cr3"
				case "epub":
					viewer = "cr3"
				case ".zip":
					viewer = "cr3"
				case ".chm":
					viewer = "xchm"
				}
				cmd := exec.Command(viewer, (bk.filepath + bk.filename))
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
fmt.Println()
}
