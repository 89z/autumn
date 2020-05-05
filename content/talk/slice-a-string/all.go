package main
func main() {
   s1 := "Sunday"
   // begin character(s)
   // a single character
   s2 := s1[:1]
   // middle character(s)
   // multiple characters
   s3 := s1[1:3]
   // end character(s)
   s4 := s1[len(s1) - 1:]
   // print
   println(s2, s3, s4)
}
