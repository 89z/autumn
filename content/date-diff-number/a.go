package main
import (
   "fmt"
   "log"
   "time"
)
func main() {
   o1, e := time.Parse(time.RFC3339[:10], "2019-12-31")
   if e != nil {
      log.Fatal(e)
   }
   o2 := time.Now()
   n := o2.Sub(o1).Seconds()
   fmt.Println(n)
}
