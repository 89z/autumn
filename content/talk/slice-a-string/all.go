package main
func main() {
   s1 := "Sunday"
   // begin character(s)
   // multiple characters
   s2 := s1[:2]
   // middle character(s)
   // a single character
   s3 := s1[1:2]
   // end character(s)
   s4 := s1[len(s1) - 1:]
   // print
   println(s2, s3, s4)
}
