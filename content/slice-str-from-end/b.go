package main
func main() {
   s_in := "♠♣♥♦"
   a_dec := []rune(s_in)
   a_tran := a_dec[len(a_dec) - 1:]
   s_enc := string(a_tran)
   println(s_enc == "♦")
}
