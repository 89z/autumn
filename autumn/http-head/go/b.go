package main

import (
   "fmt"
   "log"
   "net/http"
)

func head(s string) error {
   req, e := http.NewRequest("HEAD", s, nil)
   if e != nil {
      return e
   }
   res, e := http.DefaultClient.Do(req)
   if e != nil {
      return e
   }
   if res.StatusCode != 200 {
      return fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   return nil
}

func main() {
   e := head("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
}
