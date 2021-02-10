package main
import "path"

func main() {
   // example 1
   s1 := path.Base("C:/file.txt")
   // example 2
   s2 := path.Base(`C:\file.txt`)
   // print
   println(s1 == "file.txt", s2 == `C:\file.txt`)
}
