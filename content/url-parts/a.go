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
   // example A
   sA := o.Host
   // example B
   sB := o.Path
   // example C
   sC := o.RawQuery
   // print
   println(sA == "example.com", sB == "/one", sC == "two=even")
}
