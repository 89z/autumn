package main

import (
   "log"
   "time"
)

func main() {
   t, e := time.Parse(time.RFC3339, "2020-05-04T03:02:01Z")
   if e != nil {
      log.Fatal(e)
   }
   n := t.Unix()
   println(n == 1588561321)
}
