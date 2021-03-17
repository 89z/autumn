package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "os"
)

func GetContents(s string) ([]byte, error) {
   o, e := http.Get(s)
   if e != nil {
      return []byte{}, e
   }
   return ioutil.ReadAll(o.Body)
}

func main() {
   y, e := GetContents("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(y)
}