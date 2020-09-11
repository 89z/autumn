package main

func main() {
   n, n2 := 29, 10
   // example 1
   n3 := n / n2
   // example 2
   n4 := n % n2
   // example 3
   n5 := float64(n) / float64(n2)
   // print
   println(n3 == 2, n4 == 9, n5 == 2.9)
}
