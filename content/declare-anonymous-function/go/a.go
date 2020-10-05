package main

func main() {
   f := func (n, n1 int) bool {
      return n > n1
   }
   b := f(9, 8)
   println(b)
}
