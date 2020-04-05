package main
import (
   "fmt"
   "strconv"
)
func main() {
   var s1 = "10"
   var n1 = 1.9
   // example 1
   var n2 = int(n1)
   // example 2
   var n3, _ = strconv.ParseInt(s1, 10, 0)
   // example 3
   var n4, _ = strconv.Atoi(s1)
   // example 4
   var n5 int
   fmt.Sscan(s1, &n5)
   // print
   println(n2, n3, n4, n5)
}
