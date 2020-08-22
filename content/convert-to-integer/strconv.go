package main
import (
   "os"
   "strconv"
)
func main() {
   s := "10"
   // example 1
   n1, e := strconv.ParseInt(s, 10, 0)
   if e != nil {
      os.Exit(1)
   }
   // example 2
   n2, e := strconv.Atoi(s)
   if e != nil {
      os.Exit(1)
   }
   // print
   println(n1, n2)
}
