package main
import (
   "fmt"
   "math"
   "strconv"
)
func main() {
   // example 1
   n2 := int(n1)
   // example 2
   n3 := math.Trunc(n1)
   // example 3
   n5 := math.Round(n1)
   // example 4
   var n1, _ = strconv.ParseInt(s1, 10, 0)
   // example 5
   var n2, _ = strconv.Atoi(s1)
   // example 6
   var n3 int
   fmt.Sscan(s1, &n3)
}
