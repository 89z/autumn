package main
func main() {
   a1 := []rune("📕📙📒📗")
   a2 := a1[:2]
   s1 := string(a2)
   println(s1 == "📕📙")
}
