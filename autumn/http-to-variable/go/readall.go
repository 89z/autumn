package main

import (
   "io/ioutil"
   "net/http"
   "time"
)

func GetContents(s string) ([]byte, error) {
   o, e := http.Get(s)
   if e != nil {
      return []byte{}, e
   }
   return ioutil.ReadAll(o.Body)
}

func main() {
   GetContents("http://speedtest.lax.hivelocity.net/10Mio.dat")
   println("Sleep")
   time.Sleep(time.Duration(time.Minute))
}
