package main
import "path"

func main() {
   { // example 1
      s := path.Dir("C:/go/file.txt")
      println(s == "C:/go")
   }
   { // example 2
      s := path.Dir(`C:\go\file.txt`)
      println(s == ".")
   }
}
