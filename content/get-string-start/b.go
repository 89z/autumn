package main
func main() {
   a := []rune("📗📒📕")
   a2 := a[:2]
   s := string(a2)
   println(s == "📗📒")
}
