package main
import (
   "fmt"
   "log"
   "time"
)
func main() {
   o := time.Now()
   y, e := o.MarshalText()
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%s\n", y)
}
