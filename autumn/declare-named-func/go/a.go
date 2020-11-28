package main
// example 1
func add(n int, n1 int) int {
   return n + n1
}
// example 2
func sub(n, n2 int) int {
   return n - n2
}
// print
func main() {
   n1 := add(8, 1)
   n2 := sub(8, 1)
   println(n1 == 9, n2 == 7)
}
