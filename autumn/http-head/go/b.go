package main

import (
   "fmt"
   "net/http"
)

func head(addr string) error {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return err
   }
   res, err := new(http.Client).Do(req)
   if err != nil {
      return err
   }
   if res.StatusCode != 200 {
      return fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   return nil
}

func main() {
   err := head("http://speedtest.lax.hivelocity.net")
   if err != nil {
      panic(err)
   }
}
