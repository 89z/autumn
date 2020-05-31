package main
func main() {
   // example 1
   s1 := "Sunday"
   s2 := s1[1:2]
   // example 2
   a1 := []rune("♠♣♥♦")
   a2 := a1[1:2]
   s3 := string(a2)
   // print
   println(s2 == "u", s3 == "♣")
}
