package main

import (
   "fmt"
   "log"
   "net/url"
)

func main() {
   u, e := url.Parse("http://golang.org?west=left")
   if e != nil {
      log.Fatal(e)
   }
   q := u.Query()
   fmt.Println(q)
}
