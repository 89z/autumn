package main

import (
   "log"
   "net/http"
   "net/http/cookiejar"
   "net/url"
)

func main() {
   u := url.URL{Scheme: "http", Host: "www.deezer.com"}
   j, e := cookiejar.New(nil)
   if e != nil {
      log.Fatal(e)
   }
   http.DefaultClient.Jar = j
   http.Head(u.String())
   for _, c := range http.DefaultClient.Jar.Cookies(&u) {
      println(c.Name, c.Value)
   }
}
