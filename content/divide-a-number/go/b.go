package main
import "math"

func main() {
   n, n0 := float64(46), float64(10)
   // example 1
   n1 := math.Mod(n, n0)
   // example 2
   n2 := math.Remainder(n, n0)
   // print
   println(n1 == 6, n2 == -4)
}
