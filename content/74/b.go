package main
import "math"

func main() {
   // example 1
   n := math.Mod(29, 10)
   // example 2
   n2 := math.Remainder(29, 10)
   // print
   println(n == 9, n2 == -1)
}
