package main

import (
   "log"
   "time"
)

func main() {
   t, e := time.Parse(time.RFC3339, "2020-12-31T01:02:31Z")
   if e != nil {
      log.Fatal(e)
   }
   n := t.Unix()
   println(n == 1609376551)
}
