package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
  "io/ioutil"
//	"os/exec"
//  spl "libs_m/word"
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

func main() {

	flnumb := flag.String("n", "",   "ввод номера ")
	flname := flag.String("t", "",   "ввод названия")
	fldstart := flag.String("j", "", "ввод года ввода")
	fldend := flag.String("e", "-",   "ввод года вывода")
	flabstract := flag.String("a","", "файл аннотации")
	flkey := flag.String("k", "",    "файл ключей")
	fldocmem := flag.String("m", "-",  "примечание")
	flcountry := flag.String("c", "ru", "страна,если не 'ru'")
	flfilepath := flag.String("fp","/home/o/bibl/gost/01/","путь к файлу")
	flfilename := flag.String("fn","","имя файла ")
//	flversion := flag.String("v", "", "версия")

// прием исходных данных
	flag.Parse()
  fmt.Printf("\nrecord_add_doc: %s\n",version)
	println("  ввод номера СНиПа     -n :", *flnumb)
	println("  ввод названия         -t :", *flname)
	println("  ввод года ввода       -j :", *fldstart)
	println("  ввод года вывода      -e :", *fldend)
	println("  файл аннотации        -a :", *flabstract)
	println("  файл ключевых слов    -k :", *flkey)
	println("  примечание            -m :", *fldocmem)
	println("  страна,если не 'ru'   -c :", *flcountry)
	println("  имя файла             -fn :", *flfilename)
	println("  путь к файлу          -fp :", *flfilepath)
//	println("версия                -v :", *flversion)
	println()

	db, err := sql.Open("sqlite3", "db/snips.dbl")
	if err != nil {
		log.Fatal(err)
	}

	//  проверка указания имени файла
  if *flfilename==""{
      fmt.Println("Внимание! не указано имя файла . ")
      return
  }
  var abstract,key string
  if *flabstract!=""{
      bs,err:=ioutil.ReadFile(*flabstract)
      if err !=nil {
          fmt.Printf("нет файла аннотации: %s",*flabstract) 
          return
      }else{
          abstract=string(bs)
          abstract="'"+abstract+"'"
      }
  }
  if *flkey!=""{
      bs,err:=ioutil.ReadFile(*flkey)
      if err !=nil {
          fmt.Printf("нет файла ключевых слов: %s",*flkey) 
          return
      }else{
          key=string(bs)
          key="'"+key+"'"
      }
  }
 
 //  проверка на дублирование с БД
	//  проверка авторов

    var slct string
/*    var aut_id int64//,izd_id int
  	slct = "select * from author where name=$1;"
	  rows, err := db.Query(slct,*flauthor)
  	if err != nil {
	  	log.Fatal(err)
  	}
  	defer rows.Close()

	auts := make([]*Author, 0)
	for rows.Next() {
		au := new(Author)
		err := rows.Scan(&au.aut_id, &au.name)
		if err != nil {
			log.Fatal(err)
		}
		auts = append(auts, au)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
*/
	// печать
/*
	for _,  au:= range auts {
      fmt.Printf("Автор: id = %d\t%s\n", au.aut_id, au.name)
      if au.aut_id >0 {aut_id=au.aut_id}
	}
  // добавление автора
  if aut_id ==0 {
      fmt.Printf("\naut_id=%d\n",aut_id)
      slct="insert into author (name) values($1);"
      result, err := db.Exec(slct,*flauthor)
        if err != nil {
          log.Fatal(err)
        }
        defer rows.Close()
        aut_id,err=result.LastInsertId()
        if err !=nil {
           log.Fatal(err)
         }
        fmt.Printf("\nДобавлен автор: aut_id=%d\n",aut_id)
         
    }
*/

	//  проверка издательства
/*
//    var slct string
    var izd_id int64
  	slct = "select * from izdat where izd_name=$1;"
	  rows, err = db.Query(slct,*flizdat)
  	if err != nil {
	  	log.Fatal(err)
  	}
  	defer rows.Close()

	izds := make([]*Izdatelstwo, 0)
	for rows.Next() {
		iz := new(Izdatelstwo)
		err := rows.Scan(&iz.izd_id, &iz.izd_name)
		if err != nil {
			log.Fatal(err)
		}
		izds = append(izds, iz)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
*/
	// печать
/*
	for _,  iz:= range izds {
      fmt.Printf("Издательство: id = %d\t%s\n", iz.izd_id, iz.izd_name)
      if iz.izd_id >0 {izd_id=iz.izd_id}
	}
  // добавление издательства
  if izd_id ==0 {
      fmt.Printf("\nizd_id=%d\n",izd_id)
      slct="insert into izdat (izd_name) values($1);"
      result, err := db.Exec(slct,*flizdat)
        if err != nil {
          log.Fatal(err)
        }
        defer rows.Close()
        izd_id,err=result.LastInsertId()
        if err !=nil {
           log.Fatal(err)
         }
         fmt.Printf("\nДобавлено издательство: izd_id=%d\n",izd_id)
         
    }
*/


//  проверка книги
/*
 	slct = "select buch_id,title,size,jahr,izdan,isbn from buch where title=$1 and jahr=$2 and izdan=$3 and isbn=$4 ;"
	  rows, err = db.Query(slct,*fltitle,*fljahr,*flizdanie,*flisbn)
  	if err != nil {
	  	log.Fatal(err)
  	}
  	defer rows.Close()

	buchs := make([]*Book, 0)
	for rows.Next() {
		bu := new(Book)
		err := rows.Scan(&bu.buch_id,&bu.title, &bu.size, &bu.jahr, &bu.izdan, &bu.isbn)
		if err != nil {
			log.Fatal(err)
		}
		buchs = append(buchs, bu)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// печать
  var buch_id int64
	for _,  bu:= range buchs {
		fmt.Printf("id = %d\t%s\n%s %s %s %s\n\n", bu.buch_id, bu.title,bu.size, bu.jahr, bu.izdan, bu.isbn)
      if bu.buch_id >0 {buch_id=bu.buch_id}
	}
  */
  // добавление книги 
  //if buch_id ==0 {
  //fmt.Printf("\nbuch_id=%d\n",buch_id)
  slct="insert into doc (numb,name,abstract,dstart,dend,country,filepath,filename,doc_mem,key) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);"
  result, err := db.Exec(slct,*flnumb,*flname,abstract,*fldstart,*fldend,*flcountry,*flfilepath,*flfilename,*fldocmem,key)
  if err != nil {
     log.Fatal(err)
  }
  //defer result.Close()
  doc_id,err:=result.LastInsertId()
     if err !=nil {
        log.Fatal(err)
     }else{
        fmt.Printf("\nДобавлена запись: doc_id=%d\n",doc_id)
     }
 } //end
//	}

//подключение автора к книге (lnkaut)
// проверка наличия
/*
 	slct = "select buch_id,aut_id from lnkaut where buch_id=$1 and aut_id=$2 ;"
  row:= db.QueryRow(slct,buch_id,aut_id)
  var tmb,tma int64
  err=row.Scan(tmb,tma)
  if err == sql.ErrNoRows {
      slct="insert into lnkaut(aut_id,buch_id) values($1,$2)"
      result, err := db.Exec(slct,aut_id,buch_id)
      if err != nil {
          log.Fatal(err)
      }
      defer rows.Close()
      tmb,err=result.LastInsertId()
      if err !=nil {
           log.Fatal(err)
      }
  }
  println()
*/
// замена аннотации
/*
  if *flabstract!=""{
      bs,err:=ioutil.ReadFile(*flabstract)
      if err !=nil {
          fmt.Printf("нет файла аннотации: %s",*flabstract) 
      }else{
          str:=string(bs)
          str="'"+str+"'"
          println(buch_id)
          slct="update abstract set atext=$1 where buch_id=$2;"
          result,err:=db.Exec(slct,str,buch_id)
          if err !=nil {
               log.Fatal(err)
           }
          defer rows.Close()
          tmb,err=result.LastInsertId()
          if err !=nil {
               log.Fatal(err)
          }
          fmt.Printf("произведена замена аннотации в книге buch_id=%d\n",tmb)
      }
  }

/*
// замена оглавления 

  if *floglav!=""{
      bs,err:=ioutil.ReadFile(*floglav)
      if err !=nil {
          fmt.Printf("нет файла оглавления: %s",*floglav) 
      }else{
          str:=string(bs)
          str="'"+str+"'"
          println(buch_id)
          slct="update ogl set oglav=$1 where buch_id=$2;"
          result,err:=db.Exec(slct,str,buch_id)
          if err !=nil {
               log.Fatal(err)
           }
          defer rows.Close()
          tmb,err=result.LastInsertId()
          if err !=nil {
               log.Fatal(err)
          }
          fmt.Printf("произведена замена оглавления в книге buch_id=%d\n",tmb)
      }
  }

// замена указателя 

  if *flindex!=""{
      bs,err:=ioutil.ReadFile(*flindex)
      if err !=nil {
          fmt.Printf("нет файла указателя: %s",*flindex) 
      }else{
          str:=string(bs)
          str="'"+str+"'"
          println(buch_id)
          slct="update alf set aindex=$1 where buch_id=$2;"
          result,err:=db.Exec(slct,str,buch_id)
          if err !=nil {
               log.Fatal(err)
           }
          defer rows.Close()
          tmb,err=result.LastInsertId()
          if err !=nil {
               log.Fatal(err)
          }
          fmt.Printf("произведена замена указателя в книге buch_id=%d\n",tmb)
      }
  }
*/

//}
