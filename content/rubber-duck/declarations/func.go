package main
func main() {
   f1 := func(n1 int) int {
      return n1 + 10
   }
   n2 := f1(20)
   println(n2)
}
