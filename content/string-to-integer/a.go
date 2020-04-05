package main
import (
   "fmt"
   "strconv"
)
func main() {
   // example 4
   var n1, _ = strconv.ParseInt(s1, 10, 0)
   // example 5
   var n2, _ = strconv.Atoi(s1)
   // example 6
   var n3 int
   fmt.Sscan(s1, &n3)
}
