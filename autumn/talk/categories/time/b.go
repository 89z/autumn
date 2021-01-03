package main

import (
   "log"
   "math/big"
   "time"
)

func main() {
   t, e := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
   if e != nil {
      log.Fatal(e)
   }
   u, e := time.Parse(time.RFC3339, "2020-05-04T03:02:01Z")
   if e != nil {
      log.Fatal(e)
   }
   x := u.Sub(t).Milliseconds() / 1e3
   s := big.NewInt(x).Text(36)
   println(s == "6dv3d")
}
