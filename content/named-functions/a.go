package main
// example 1
func f1(n int, n1 int) bool {
   return n > n1
}
// example 2
func f2(n, n2 int) bool {
   return n > n2
}
// print
func main() {
   b1 := f1(9, 8)
   b2 := f2(9, 8)
   println(b1, b2)
}
