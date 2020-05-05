package main
func main() {
   s1 := "♠"
   // example 1
   n1 := len(s1)
   // example 2
   a1 := []rune(s1)
   n2 := len(a1)
   // print
   println(n1 == 3, n2 == 1)
}
