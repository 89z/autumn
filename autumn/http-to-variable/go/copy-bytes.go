package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "time"
)

func GetContents(s string) (bytes.Buffer, error) {
   buf_o := bytes.Buffer{}
   get_o, e := http.Get(s)
   if e != nil {
      return buf_o, e
   }
   n, e := io.Copy(&buf_o, get_o.Body)
   if e != nil {
      return buf_o, fmt.Errorf("%v %v", n, e)
   }
   return buf_o, nil
}

func main() {
   GetContents("http://speedtest.lax.hivelocity.net/10Mio.dat")
   println("Sleep")
   time.Sleep(time.Duration(time.Minute))
}
