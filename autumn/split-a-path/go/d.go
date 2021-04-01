package main
import "path/filepath"

func main() {
   { // example 1
      s := filepath.Dir("C:/go/file.txt")
      println(s == `C:\go`)
   }
   { // example 2
      s := filepath.Dir(`C:\go\file.txt`)
      println(s == `C:\go`)
   }
}
