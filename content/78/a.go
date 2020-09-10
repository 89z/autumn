package main
// example 1
func f(s string) int {
   return len(s)
}
// example 2
var f2 = func(s string) int {
   return len(s)
}
// print
func main() {
   n := f("May")
   n2 := f2("May")
   println(n == 3, n2 == 3)
}
