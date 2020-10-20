package main

func main() {
   // example 1
   f1 := func (n int, n1 int) bool {
      return n > n1
   }
   // example 2
   f2 := func (n, n2 int) bool {
      return n > n2
   }
   // print
   b1 := f1(9, 8)
   b2 := f2(9, 8)
   println(b1, b2)
}
