package main
import "fmt"
// example 1
func f1(n1 int) int {
   return n1 + 10
}
// example 2
var f2 = func(n1 int) int {
   return n1 + 10
}
// print
func main() {
   fmt.Printf("%T\n", f1)
   fmt.Printf("%T\n", f2)
}
