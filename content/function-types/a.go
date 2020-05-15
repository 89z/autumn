package main
// example 1
func f1(s1 string) int {
   return len(s1)
}
// example 2
var f2 = func(s1 string) int {
   return len(s1)
}
// print
func main() {
   n1 := f1("Sunday")
   n2 := f2("Sunday")
   println(n1, n2)
}
