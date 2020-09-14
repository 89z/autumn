package main
import "math"

func main() {
   n, n2 := float64(46), float64(10)
   // example 1
   n3 := math.Mod(n, n2)
   // example 2
   n4 := math.Remainder(n, n2)
   // print
   println(n3 == 6, n4 == -4)
}
