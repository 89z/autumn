package main

import (
   "os"
   "time"
)

func main() {
   s_layout := time.RFC3339[:10]
   o, e := time.Parse(s_layout, "2019-12-31")
   if e != nil {
      os.Exit(1)
   }
   o2 := time.Now()
   n := o2.Sub(o).Seconds()
   println(n)
}
