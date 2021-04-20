package main
import "strings"

func main() {
   {
      b := strings.HasPrefix("north", "no")
      println(b) // true
   }
   {
      b := strings.HasPrefix("north", "")
      println(b) // true
   }
   {
      b := strings.HasPrefix("", "")
      println(b) // true
   }
   {
      b := strings.HasPrefix("", "north")
      println(b) // false
   }
}
