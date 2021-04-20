package main
import "strings"

func main() {
   {
      b := strings.Contains("", "")
      println(b) // true
   }
   {
      b := strings.Contains("north", "")
      println(b) // true
   }
   {
      b := strings.Contains("north", "or")
      println(b) // true
   }
}
