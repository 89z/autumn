package main
import (
   "log"
   "strconv"
)
func main() {
   s := "1.9"
   n, e := strconv.ParseFloat(s, 64)
   if e != nil {
      log.Fatal(e)
   }
   println(n)
}
