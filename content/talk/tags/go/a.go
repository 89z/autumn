package main
import (
   "log"
   "strconv"
)
func main() {
   n, e := strconv.Atoi("10")
   if e != nil {
      log.Fatal(e)
   }
   println(n)
}
