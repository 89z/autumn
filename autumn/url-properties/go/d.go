package main

import (
   "log"
   "net/url"
)

func main() {
   v, e := url.ParseQuery("month=March&day=Friday")
   if e != nil {
      log.Fatal(e)
   }
   s := v.Get("day")
   println(s)
}
