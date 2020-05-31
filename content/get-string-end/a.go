package main
func main() {
   s1 := "Sunday"
   // example 1
   s2 := s1[len(s1) - 1:]
   // example 2
   s3 := s1[len(s1) - 2:]
   // print
   println(s2 == "y", s3 == "ay")
}
