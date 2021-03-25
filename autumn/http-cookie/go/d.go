package main

import (
   "log"
   "net/http"
   "net/http/cookiejar"
)

func main() {
   r, e := http.NewRequest("GET", "http://www.deezer.com", nil)
   if e != nil {
      log.Fatal(e)
   }
   c, e := cookiejar.New(nil)
   if e != nil {
      log.Fatal(e)
   }
   h := http.Client{Jar: c}
   h.Do(r)
   for _, d := range c.Cookies(r.URL) {
      println(d.Name, d.Value)
   }
}
