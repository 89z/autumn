package main

import (
   "log"
   "net/http"
   "os"
)

func main() {
   in_o, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   out_o, e := os.Create("index.html")
   if e != nil {
      log.Fatal(e)
   }
   out_o.ReadFrom(in_o.Body)
}
