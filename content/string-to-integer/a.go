package main
import (
   "fmt"
   "strconv"
)
func main() {
   // example 2
   n2, _ := strconv.ParseInt("11", 10, 0)
   // example 3
   n3, _ := strconv.Atoi("11")
   // example 4
   var n4 int
   fmt.Sscan("11", &n4)
}
