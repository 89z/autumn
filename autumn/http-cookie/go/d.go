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
   j.SetCookies(r.URL, []*http.Cookie{
      {Name: "month", Value: "March"},
   })
}
