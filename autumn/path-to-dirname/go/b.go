package main
import "path/filepath"

func main() {
   // example 1
   s1 := filepath.Dir("C:/go/file.txt")
   // example 2
   s2 := filepath.Dir(`C:\go\file.txt`)
   // print
   println(s1 == `C:\go`, s2 == `C:\go`)
}
