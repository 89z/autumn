package main

import (
   "io"
   "net/http"
   "os"
   "strings"
)

func main() {
   o_in, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      os.Exit(1)
   }
   o_out := strings.Builder{}
   io.Copy(&o_out, o_in.Body)
   s := o_out.String()
   print(s)
}
