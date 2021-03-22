package main
import "path/filepath"

func main() {
   { // example 1
      s := filepath.Base("C:/file.txt")
      println(s == "file.txt")
   }
   { // example 2
      s := filepath.Base(`C:\file.txt`)
      println(s == "file.txt")
   }
}
