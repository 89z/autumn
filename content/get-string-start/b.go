package main

func main() {
   a := []rune("📗📒📕")
   a1 := a[:2]
   s := string(a1)
   println(s == "📗📒")
}
