package main
// example 1
func f1(n int, n1 int) int {
   return n + n1
}
// example 2
func f2(n, n2 int) int {
   return n + n2
}
// print
func main() {
   n1 := f1(1, 2)
   n2 := f2(1, 2)
   println(n1 == 3, n2 == 3)
}
