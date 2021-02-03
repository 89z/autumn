package main

import (
   "fmt"
   "log"
   "net/url"
)

func main() {
   s := "http://golang.org"
   u, e := url.Parse(s)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(u)
}
