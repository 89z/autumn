package main

import (
   "fmt"
   "log"
   "time"
)

func main() {
   s := time.RFC3339[:19]
   o, e := time.Parse(s, "2019-12-31T00:00:00")
   if e != nil {
      log.Fatal(e)
   }
   o1a, e := time.Parse(s, "2019-12-31T23:59:59")
   if e != nil {
      log.Fatal(e)
   }
   o1 := o1a.Sub(o)
   fmt.Println(o1)
}
