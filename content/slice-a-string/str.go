package main
func main() {
   s1 := "Sunday"
   // example 1
   s2 := s1[:1]
   // example 2
   s3 := s1[1:2]
   // example 3
   s4 := s1[1:]
   // example 4
   s5 := s1[len(s1) - 1:]
   // print
   println(s2 == "S", s3 == "u", s4 == "unday", s5 == "y")
}
