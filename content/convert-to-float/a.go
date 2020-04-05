package main
import (
   "fmt"
   "strconv"
)
func main() {
   var s1 = "1.1"
   // example 1
   var n1, _ = strconv.ParseFloat(s1, 0)
   // example 2
   var n2 float64
   fmt.Sscan(s1, &n2)
   // print
   fmt.Println(n1, n2)
}
