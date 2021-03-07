package main

import (
   "log"
   "net/http"
   "net/url"
   "os"
)

func main() {
   req := &http.Request{
      URL: &url.URL{Scheme: "http", Host: "speedtest.lax.hivelocity.net"},
   }
   resp, e := http.DefaultClient.Do(req)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.ReadFrom(resp.Body)
}
