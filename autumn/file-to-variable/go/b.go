package main

import (
   "bytes"
   "log"
   "net/http"
)

func main() {
   o_in, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   o_out := bytes.Buffer{}
   n, e := o_out.ReadFrom(o_in.Body)
   if e != nil {
      log.Fatal(e)
   }
   print(n, o_out.String())
}
