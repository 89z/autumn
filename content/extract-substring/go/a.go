package main

func main() {
   s := "May"
   // example 1
   s1 := s[0]
   // example 2
   s2 := s[:1]
   // example 3
   s3 := s[1:2]
   // example 4
   s4 := s[1:]
   // print
   println(s1 == 'M', s2 == "M", s3 == "a", s4 == "ay")
}
