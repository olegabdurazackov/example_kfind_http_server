// для вил
//
package main
import (
    "fmt"
    "io/ioutil"
    "os"
//    "strings"
//    "sort"
    spl "libs_m/word"
//    "regexp"
)
func main()  {
    if len(os.Args)!=2{
        println("Usage:")
        println("   zoomout filenamein > filenameout")
        return
    }
    infile:=os.Args[1]
    fmt.Printf("\n zoomout v0-02 create zoom out file : %s\n",infile)
    bs,err:=ioutil.ReadFile(infile)
    if err !=nil {
        fmt.Println(err)
         return 
     }
     str:=string(bs)
     z:=make([]string,0)
     spl.Split_to_word(&str,&z)
     j:=len(z)
     fmt.Println(j)
     for i:=0;i<j;i++{
        fmt.Printf( "%s ",z[i]) //("\nx(%d)=%s\n",i,x[i])
    }
  return 
}
