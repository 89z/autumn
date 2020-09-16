package main

func main() {
   f := func (n, n2 int) bool {
      return n > n2
   }
   b := f(9, 8)
   println(b)
}
