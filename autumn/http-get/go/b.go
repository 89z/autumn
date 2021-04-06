package main

import (
   "net/http"
   "os"
)

func get(s string) (*http.Response, error) {
   r, e := http.NewRequest("GET", s, nil)
   if e != nil { return nil, e }
   return http.DefaultClient.Do(r)
}

func main() {
   r, e := get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      panic(e)
   }
   os.Stdout.ReadFrom(r.Body)
}
