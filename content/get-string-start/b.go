package main
func main() {
   a := []rune("♠♣♥♦")
   a2 := a[:1]
   s := string(a2)
   println(s == "♠")
}
