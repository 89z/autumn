package main
import (
   "fmt"
   "strconv"
)
func main() {
   var s1 = "1.9"
   var n1 = 10
   // example 1
   var n2 = float64(n1)
   // example 2
   var n3, _ = strconv.ParseFloat(s1, 0)
   // example 3
   var n4 float64
   fmt.Sscan(s1, &n4)
   // print
   println(n2, n3, n4)
}
