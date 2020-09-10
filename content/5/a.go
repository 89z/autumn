package main
import (
   "os"
   "time"
)
func main() {
   o1, e := time.Parse(time.RFC3339[:10], "2019-12-31")
   if e != nil {
      os.Exit(1)
   }
   o2 := time.Now()
   n := o2.Sub(o1).Seconds()
   println(n)
}
