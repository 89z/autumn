package main
import "math"
func main() {
   n1 := 1.9
   // example 1
   n2 := int(n1)
   // example 2
   n3 := math.Trunc(n1)
   // example 3
   n4 := math.Round(n1)
   // print
   println(n2, n3, n4)
}
