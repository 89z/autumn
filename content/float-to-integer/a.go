package main
import "math"
func main() {
   n1 := -8.9
   // truncate
   n2 := int(n1)
   // truncate
   n3 := math.Trunc(n1)
   // round
   n4 := int(n1 + math.Copysign(.5, n1))
   // round
   n5 := math.Round(n1)
   // print
   print(n2, "\n", n3, "\n", n4, "\n", n5, "\n")
}
