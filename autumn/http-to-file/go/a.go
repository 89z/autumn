package main

import (
   "io"
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
   return io.Copy(create, get.Body)
}

func main() {
   httpCopy("http://speedtest.lax.hivelocity.net", "index.html")
}
