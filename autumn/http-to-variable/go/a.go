package main

import (
   "bytes"
   "log"
   "net/http"
)

func main() {
   get_o, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   buf_o := bytes.Buffer{}
   out_o.ReadFrom(in_o.Body)
   print(out_o.String())
}
