package main

import (
   "net/http"
   "os"
)

func main() {
   r, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      panic(e)
   }
   defer r.Body.Close()
   os.Stdout.ReadFrom(r.Body)
}
