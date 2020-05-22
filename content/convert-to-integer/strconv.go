package main
import (
   "log"
   "strconv"
)
func main() {
   s1 := "10"
   // example 1
   n1, e := strconv.ParseInt(s1, 10, 0)
   if e != nil {
      log.Fatal(e)
   }
   // example 2
   n2, e := strconv.Atoi(s1)
   if e != nil {
      log.Fatal(e)
   }
   // print
   println(n1, n2)
}
