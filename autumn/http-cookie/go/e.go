package main

import (
   "log"
   "net/http"
   "net/http/cookiejar"
)

func main() {
   r, e := http.NewRequest("GET", "http://deezer.com", nil)
   if e != nil {
      log.Fatal(e)
   }
   c, e := cookiejar.New(nil)
   if e != nil {
      log.Fatal(e)
   }
   c.SetCookies(r.URL, []*http.Cookie{
      {Name: "west", Value: "right"}
   })
}
