package main
import "fmt"
func main() {
   s1 := "sunday"
   // example 1
   s2 := s1[3:4]
   // example 2
   s3 := s1[3:]
   // example 3
   s4 := s1[:3]
   // example 4
   s5 := s1[len(s1) - 3:]
   // example 5
   s6 := []rune("☺ U+263A")[:1]
   // print
   fmt.Printf("%q\n", s2)
   fmt.Printf("%q\n", s3)
   fmt.Printf("%q\n", s4)
   fmt.Printf("%q\n", s5)
   fmt.Printf("%q\n", s6)
}
