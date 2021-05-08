package main

import (
   "fmt"
   "net/http"
)

func main() {
   {
      r, e := http.Head("http://speedtest.lax.hivelocity.net")
      println(r.StatusCode == 200, e == nil)
   }
   {
      r, e := http.Head("http://speedtest.lax.hivelocity.net/404")
      println(r.StatusCode == 404, e == nil)
   }
}
