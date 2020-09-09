package main

func main() {
   // example 1
   s := "June"
   s2 := s[1:3]
   // example 2
   a := []rune("📗📒📙📕")
   a2 := a[1:3]
   s3 := string(a2)
   // print
   println(s2 == "un", s3 == "📒📙")
}
