package main

import (
   "bytes"
   "fmt"
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
   n, e := buf_o.ReadFrom(get_o.Body)
   if e != nil {
      return buf_o, fmt.Errorf("%v %v", n, e)
   }
   return buf_o, nil
}

func main() {
   o, e := GetContents("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   o.WriteTo(os.Stdout)
}
