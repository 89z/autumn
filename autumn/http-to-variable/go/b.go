package main

import (
   "io"
   "net/http"
   "os"
   "strings"
)

func main() {
   in_o, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      os.Exit(1)
   }
   out_o := strings.Builder{}
   io.Copy(&out_o, in_o.Body)
   s := out_o.String()
   print(s)
}
