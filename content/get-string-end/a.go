package main
func main() {
   s := "Sunday"
   // example 1
   s1 := s[1:2]
   // example 2
   s2 := s[1:]
   // example 3
   s3 := s[len(s) - 1:]
   // print
   println(s1 == "u", s2 == "unday", s3 == "y")
}
