package main
import "unicode/utf8"
func main() {
   s_in := "♠♣♥♦"
   a_tran, _ := utf8.DecodeRuneInString(s_in)
   println(string(a_tran))
}