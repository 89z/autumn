package main
func main() {
   // example 1
   s1 := "Sunday"
   s2 := s1[1:3]
   // example 2
   a1 := []rune("🌕🌗🌑🌓")
   a2 := a1[1:3]
   s3 := string(a2)
   // print
   println(s2 == "un", s3 == "🌗🌑")
}
