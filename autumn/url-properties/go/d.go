package main

import (
   "log"
   "net/url"
)

func main() {
   q, e := url.ParseQuery("west=left&east=right")
   if e != nil {
      log.Fatal(e)
   }
   s := q.Get("west")
   println(s == "left")
}
