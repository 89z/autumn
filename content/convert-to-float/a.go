package main
import (
   "fmt"
   "strconv"
)
func main() {
   // example 1
   var n1 = 10
   var n2 = float64(n1)
   // example 2
   var s1 = "1.9"
   var n3, _ = strconv.ParseFloat(s1, 0)
   // example 3
   var s2 = "1.9"
   var n4 float64
   fmt.Sscan(s2, &n4)
   // print
   println(n2, n3, n4)
}
