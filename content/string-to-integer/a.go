package main
import (
   "fmt"
   "strconv"
)
func main() {
   var s1 = "10"
   // example 1
   var n1, _ = strconv.ParseInt(s1, 10, 0)
   // example 2
   var n2, _ = strconv.Atoi(s1)
   // example 3
   var n3 int
   fmt.Sscan(s1, &n3)
   // print
   fmt.Println(n1, n2, n3)
}
