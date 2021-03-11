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
   j, e := cookiejar.New(nil)
   if e != nil {
      log.Fatal(e)
   }
   c := http.Client{Jar: j}
   c.Do(r)
   for _, d := range j.Cookies(r.URL) {
      println(d.Name, d.Value)
   }
}
