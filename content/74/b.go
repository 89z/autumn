package main
import "math"
func main() {
   // example 1
   n1 := math.Mod(29, 10)
   // example 2
   n2 := math.Remainder(29, 10)
   // print
   println(n1 == 9, n2 == -1)
}
