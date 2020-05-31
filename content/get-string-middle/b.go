package main
func main() {
   a := []rune("♠♣♥♦")
   a2 := a[1:2]
   s2 := string(a2)
   println(s2 == "♣")
}
