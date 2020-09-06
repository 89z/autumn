package main
func main() {
   n1 := 29
   n2 := 10
   // example 1
   n3 := n1 / n2
   // example 2
   n4 := n1 % n2
   // example 3
   n5 := float64(n1) / float64(n2)
   // print
   println(n3 == 2, n4 == 9, n5 == 2.9)
}
