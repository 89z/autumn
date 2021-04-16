package main

import (
   "net/http"
   "os"
)

func main() {
   req, e := http.NewRequest("GET", "http://speedtest.lax.hivelocity.net", nil)
   if e != nil {
      panic(e)
   }
   res, e := new(http.Client).Do(req)
   if e != nil {
      panic(e)
   }
   os.Stdout.ReadFrom(res.Body)
}
