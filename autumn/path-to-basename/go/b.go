package main
import "path/filepath"

func main() {
   // example 1
   s1 := filepath.Base("C:/file.txt")
   // example 2
   s2 := filepath.Base(`C:\file.txt`)
   // print
   println(s1 == "file.txt", s2 == "file.txt")
}
