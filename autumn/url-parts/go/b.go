package main

import (
   "log"
   "net/url"
)

func main() {
   u, e := url.Parse("https://example.com/one?two=even")
   if e != nil {
      log.Fatal(e)
   }
   s := u.Hostname()
   println(s == "example.com")
}
