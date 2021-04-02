package main

import (
   "log"
   "net/http"
   "os"
)

func main() {
   res, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.ReadFrom(res.Body)
}
