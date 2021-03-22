package main
import "path"

func main() {
   { // example 1
      s := path.Base("C:/file.txt")
      println(s == "file.txt")
   }
   { // example 2
      s := path.Base(`C:\file.txt`)
      println(s == `C:\file.txt`)
   }
}
