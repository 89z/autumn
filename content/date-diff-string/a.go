package main
import (
   "fmt"
   "log"
   "time"
)
func f_parse(s_val string) time.Time {
   s_lay := time.RFC3339[:19]
   o, e := time.Parse(s_lay, s_val)
   if e != nil {
      log.Fatal(e)
   }
   return o
}
func main() {
   o1 := f_parse("2019-12-31T00:00:00")
   o2 := f_parse("2019-12-31T23:59:59")
   o3 := o2.Sub(o1)
   fmt.Println(o3)
}
