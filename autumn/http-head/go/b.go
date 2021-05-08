package main

import (
   "fmt"
   "net/http"
)

func head(addr string) error {
   req, e := http.NewRequest("HEAD", addr, nil)
   if e != nil { return e }
   res, e := new(http.Client).Do(req)
   if e != nil { return e }
   if res.StatusCode != 200 {
      return fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   return nil
}

func main() {
   e := head("http://speedtest.lax.hivelocity.net")
   if e != nil {
      panic(e)
   }
}
