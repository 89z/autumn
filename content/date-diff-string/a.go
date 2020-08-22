package main
import (
   "fmt"
   "os"
   "time"
)
func main() {
   s := time.RFC3339[:19]
   o1, e := time.Parse(s, "2019-12-31T00:00:00")
   if e != nil {
      os.Exit(1)
   }
   o2, e := time.Parse(s, "2019-12-31T23:59:59")
   if e != nil {
      os.Exit(1)
   }
   o3 := o2.Sub(o1)
   fmt.Println(o3)
}
