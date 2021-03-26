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
   u, e := time.Parse(time.RFC3339, "2020-12-31T01:02:31Z")
   if e != nil {
      log.Fatal(e)
   }
   n := u.Sub(t).Milliseconds()
   s := big.NewInt(n).Text(62)
   println(s == "yqtA1i")
}
