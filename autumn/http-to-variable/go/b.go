package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "os"
)

func main() {
   get, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   b, e := ioutil.ReadAll(get.Body)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(b)
}
