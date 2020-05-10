package main
import "strings"
func main() {
   s1 := " ab "
   // example 1
   s2 := strings.Trim(s1, " ")
   // example 2
   s3 := strings.TrimLeft(s1, " ")
   // example 3
   s4 := strings.TrimRight(s1, " ")
   // example 4
   s5 := strings.TrimSpace(s1)
   // print
   println(s2 == "ab", s3 == "ab ", s4 == " ab", s5 == "ab")
}
