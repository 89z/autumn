package main

import (
   "log"
   "net/url"
)

func main() {
   o, e := url.Parse("https://example.com/one?two=even")
   if e != nil {
      log.Fatal(e)
   }
   m := o.Query()
   log.Print(m)
}