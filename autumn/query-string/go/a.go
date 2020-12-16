package main

import (
   "fmt"
   "log"
   "net/url"
)

func main() {
   m, e := url.ParseQuery("one=odd&two=even")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
