package main

import (
   "fmt"
   "net/url"
)

func main() {
   u, e := url.Parse("http://golang.org?west=left")
   if e != nil {
      panic(e)
   }
   q := u.Query()
   fmt.Println(q)
}
