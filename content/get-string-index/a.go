package main
import "strings"
func main() {
   s1 := "Sun Mon"
   // example 1
   n1 := strings.Index(s1, "n")
   // example 2
   n2 := strings.LastIndex(s1, "n")
   // print
   println(n1, n2)
}
