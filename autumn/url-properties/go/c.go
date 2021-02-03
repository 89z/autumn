package main

import (
   "fmt"
   "log"
   "net/url"
)

func main() {
   u, e := url.Parse("http://golang.org?month=May&day=Friday")
   if e != nil {
      log.Fatal(e)
   }
   v := u.Query()
   fmt.Println(v)
}
