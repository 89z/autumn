package main

import (
   "log"
   "net/http"
   "os"
)

func main() {
   get_o, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   create_o, e := os.Create("index.html")
   if e != nil {
      log.Fatal(e)
   }
   create_o.ReadFrom(get_o.Body)
}
