package main

import (
   "log"
   "net/http"
)

func main() {
   r, e := http.NewRequest("GET", "http://www.deezer.com", nil)
   if e != nil {
      log.Fatal(e)
   }
   r.Header.Set("Cookie", "month=March")
}
