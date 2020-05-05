package main
import "strings"
func main() {
   s1 := "Sunday"
   // example 1
   b1 := strings.HasPrefix(s1, "Su")
   // example 2
   b2 := strings.Contains(s1, "un")
   // example 3
   b3 := strings.HasSuffix(s1, "ay")
   // print
   println(b1, b2, b3)
}
