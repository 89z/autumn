package main

import (
   "fmt"
   "net/http"
)

func head(s string) error {
   req, e := http.NewRequest("HEAD", s, nil)
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
