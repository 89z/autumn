package main
import "strings"

func main() {
   s := "March"
   // example 1
   b1 := strings.Contains(s, "ar")
   // example 2
   b2 := strings.Contains(s, "")
   // print
   println(b1, b2)
}
