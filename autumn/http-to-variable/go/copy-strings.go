package main

import (
   "fmt"
   "io"
   "net/http"
   "strings"
   "time"
)

func GetContents(s string) (strings.Builder, error) {
   build_o := strings.Builder{}
   get_o, e := http.Get(s)
   if e != nil {
      return build_o, e
   }
   n, e := io.Copy(&build_o, get_o.Body)
   if e != nil {
      return build_o, fmt.Errorf("%v %v", n, e)
   }
   return build_o, nil
}

func main() {
   GetContents("http://speedtest.lax.hivelocity.net/10Mio.dat")
   println("Sleep")
   time.Sleep(time.Duration(time.Minute))
}
