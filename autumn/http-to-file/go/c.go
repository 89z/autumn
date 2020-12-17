package main

import (
   "net/http"
   "os"
)

func Copy(source, dest string) (int64, error) {
   get_o, e := http.Get(source)
   if e != nil {
      return 0, e
   }
   create_o, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   return create_o.ReadFrom(get_o.Body)
}

func main() {
   Copy("http://speedtest.lax.hivelocity.net", "index.html")
}
