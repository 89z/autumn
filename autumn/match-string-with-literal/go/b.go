package main
import "strings"

func main() {
   s := "north"
   { // example 1
      b := strings.Contains(s, "or")
      println(b)
   }
   { // example 2
      b := strings.Contains(s, "")
      println(b)
   }
}
