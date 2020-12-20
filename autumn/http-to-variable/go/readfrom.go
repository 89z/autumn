package main

import (
   "bytes"
   "log"
   "net/http"
   "os"
)

func GetContents(s string) (bytes.Buffer, error) {
   buf_o := bytes.Buffer{}
   get_o, e := http.Get(s)
   if e != nil {
      return buf_o, e
   }
   buf_o.ReadFrom(get_o.Body)
   return buf_o, nil
}

func main() {
   o, e := GetContents("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   o.WriteTo(os.Stdout)
}
