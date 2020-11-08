package main

func main() {
   n, n0 := 46, 10
   // example 1
   n1 := n / n0
   // example 2
   n2 := n % n0
   // example 3
   n3 := float64(n) / float64(n0)
   // print
   println(n1 == 4, n2 == 6, n3 == 4.6)
}
