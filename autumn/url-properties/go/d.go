package main

import (
   "log"
   "net/url"
)

func main() {
   q, e := url.ParseQuery("month=March&day=Friday")
   if e != nil {
      log.Fatal(e)
   }
   s := q.Get("day")
   println(s)
}
