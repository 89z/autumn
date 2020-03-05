package main
import (
   "fmt"
   "strconv"
)
func main() {
   // example 1
   n1, _ := strconv.ParseInt("012", 10, 0)
   // example 2
   n2, _ := strconv.ParseFloat("01.2", 0)
   // print
   fmt.Println(n1, n2)
}
