package main
import "strings"
func main() {
   s1 := "Sunday"
   // example 1
   b1 := strings.Contains(s1, "und")
   // example 2
   b2 := strings.HasPrefix(s1, "Sun")
   // example 3
   b3 := strings.HasSuffix(s1, "day")
   // print
   println(b1, b2, b3)
}
