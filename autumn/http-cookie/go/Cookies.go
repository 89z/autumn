package main

import (
   "log"
   "net/http"
)

func main() {
   r, e := http.Get("https://stackoverflow.com")
   if e != nil {
      log.Fatal(e)
   }
   for _, c := range r.Cookies() {
      println(c.Name, c.Value)
   }
}
