package main

import (
   "bytes"
   "log"
   "net/http"
)

func main() {
   in_o, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   out_o := bytes.Buffer{}
   n, e := out_o.ReadFrom(in_o.Body)
   if e != nil {
      log.Fatal(e)
   }
   print(n, out_o.String())
}
