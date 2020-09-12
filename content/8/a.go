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
   n := o.Unix()
   println(n)
}
