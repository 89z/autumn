package main
import "unicode/utf8"
func main() {
   s1 := "♠"
   n1 := utf8.RuneCountInString(s1)
   println(n1 == 1)
}
