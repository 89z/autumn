package main
import (
   "os"
   "strconv"
)
func main() {
   n, e := strconv.ParseInt("11", 10, 0)
   if e != nil {
      os.Exit(1)
   }
   println(n)
}
