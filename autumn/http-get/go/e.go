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
   res, e := new(http.Transport).RoundTrip(req)
   if e != nil {
      panic(e)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
