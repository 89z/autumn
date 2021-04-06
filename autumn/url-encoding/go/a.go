package main

import (
   "fmt"
   "net/url"
)

func main() {
   s := "http://golang.org"
   u, e := url.Parse(s)
   if e != nil {
      panic(e)
   }
   fmt.Println(u)
}
