package main
import "strings"
func main() {
   s1 := "sun mon tue"
   // example 1
   n1 := strings.Index(s1, " ")
   // example 2
   n2 := strings.LastIndex(s1, " ")
   // print
   println(n1, n2)
}
