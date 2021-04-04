package main

import (
   "net/http"
   "log"
)

func main() {
   r, e := http.Get("https://stackoverflow.com")
   if e != nil {
      log.Fatal(e)
   }
   s := r.Header.Get("Set-Cookie")
   println(s)
}
