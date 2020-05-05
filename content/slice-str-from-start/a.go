package main
func main() {
   s1 := "Sunday"
   // example 1
   s2 := s1[0:1]
   // example 2
   s3 := s1[:1]
   // example 3
   s4 := s1[0:2]
   // example 4
   s5 := s1[:2]
   // print
   println(s2 == "S", s3 == "S", s4 == "Su", s5 == "Su")
}
