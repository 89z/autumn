package main
import "math"
func main() {
   // example 1
   n1 := math.Mod(19, 10)
   // example 2
   n2 := math.Remainder(19, 10)
   // print
   println(n1 == 9, n2 == -1)
}
