package main

import (
   "net/url"
   "os"
)

func main() {
   o, e := url.Parse("https://example.com/one?two=even")
   if e != nil {
      os.Exit(1)
   }
   s := o.Hostname()
   println(s == "example.com")
}