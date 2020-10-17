package main

func main() {
   s := "March"
   // example 1
   s1 := s[2]
   // example 2
   s2 := s[:1]
   // example 3
   s3 := s[2:3]
   // example 4
   s4 := s[2:]
   // print
   println(s1 == 'r', s2 == "M", s3 == "r", s4 == "rch")
}
