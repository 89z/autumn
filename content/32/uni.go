package main
import "unicode/utf8"

func main() {
   s := "📗"
   n := utf8.RuneCountInString(s)
   println(n == 1)
}
