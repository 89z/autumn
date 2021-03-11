package main

import (
   "log"
   "net/http"
)

func main() {
   r, e := http.Head("http://www.deezer.com")
   if e != nil {
      log.Fatal(e)
   }
   for _, c := range r.Cookies() {
      println(c.Name, c.Value)
   }
}
