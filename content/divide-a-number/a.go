package main

func main() {
   n, n2 := 46, 10
   // example 1
   n3 := n / n2
   // example 2
   n4 := n % n2
   // example 3
   n5 := float64(n) / float64(n2)
   // print
   println(n3 == 4, n4 == 6, n5 == 4.6)
}
