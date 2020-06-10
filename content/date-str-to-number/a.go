package main
import (
   "log"
   "time"
)
func main() {
   s := "2019-12-31"
   o, e := time.Parse(time.RFC3339[:10], s)
   if e != nil {
      log.Fatal(e)
   }
   n := o.Unix()
   println(n)
}
