package main

func main() {
   // example 1
   s := "June"
   s1 := s[1:3]
   // example 2
   a := []rune("📗📒📙📕")
   a2 := a[1:3]
   s2 := string(a2)
   // print
   println(s1 == "un", s2 == "📒📙")
}
