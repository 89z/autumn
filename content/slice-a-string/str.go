package main
func main() {
   s1 := "Sunday"
   // example 1
   s2 := s1[:3]
   // example 2
   s3 := s1[3:]
   // example 3
   s4 := s1[len(s1) - 3:]
   // print
   println(s2 == "Sun", s3 == "day", s5 == "day")
}
