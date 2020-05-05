package main
func main() {
   s1 := "Sunday"
   // start one
   s2 := s1[:1]
   // start two
   s3 := s1[:2]
   // mid one
   s4 := s1[1:2]
   // end one
   s5 := s1[len(s1) - 1:]
   // end two
   s6 := s1[len(s1) - 2:]
   // print
   println(s2 == "S", s3 == "Su", s4 == "u", s5 == "y", s6 == "ay")
}
