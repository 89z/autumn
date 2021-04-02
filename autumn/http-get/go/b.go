package main

import (
   "log"
   "net/http"
   "os"
)

func main() {
   req, e := http.NewRequest("GET", "http://speedtest.lax.hivelocity.net", nil)
   if e != nil {
      log.Fatal(e)
   }
   res, e := http.DefaultClient.Do(req)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.ReadFrom(res.Body)
}
