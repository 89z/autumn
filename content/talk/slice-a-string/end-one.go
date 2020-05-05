package main
func main() {
   s1 := "Sunday"
   s2 := s1[-1:]
   println(s2) // invalid slice index -1 (index must be non-negative)
}
