package main
import "strings"

func main() {
   s := "March"
   { // example 1
      b := strings.Contains(s, "ar")
      println(b)
   }
   { // example 2
      b := strings.Contains(s, "")
      println(b)
   }
}
