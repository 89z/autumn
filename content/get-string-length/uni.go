package main
import "unicode/utf8"
func main() {
   n := utf8.RuneCountInString("📕")
   println(n == 1)
}
