package main

import (
   "net/http"
   "os"
)

func httpCopy(source, dest string) (int64, error) {
   get, e := http.Get(source)
   if e != nil {
      return 0, e
   }
   create, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   defer create.Close()
   return create.ReadFrom(get.Body)
}

func main() {
   httpCopy("http://speedtest.lax.hivelocity.net", "index.html")
}
