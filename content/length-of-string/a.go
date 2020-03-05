package main
import "unicode/utf8"
func main() {
   s1 := "☺"
   // bytes
   n1 := len(s1)
   // characters
   n2 := utf8.RuneCountInString(s1)
   // print
   println(n1, n2)
}
