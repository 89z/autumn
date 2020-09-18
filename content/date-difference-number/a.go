package main

import (
   "os"
   "time"
)

func main() {
   s := time.RFC3339[:10]
   o, e := time.Parse(s, "2019-12-31")
   if e != nil {
      os.Exit(1)
   }
   o1 := time.Now()
   n := o1.Sub(o).Seconds()
   println(n)
}
