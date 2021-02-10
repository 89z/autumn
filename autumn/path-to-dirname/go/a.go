package main
import "path"

func main() {
   // example 1
   s1 := path.Dir("C:/go/file.txt")
   // example 2
   s2 := path.Dir(`C:\go\file.txt`)
   // print
   println(s1 == "C:/go", s2 == ".")
}
