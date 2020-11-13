package main

import (
   "io"
   "log"
   "net/http"
   "os"
)

func main() {
   in_o, e := http.Get("https://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   out_o, e := os.Create("index.html")
   if e != nil {
      log.Fatal(e)
   }
   io.Copy(out_o, in_o.Body)
}
