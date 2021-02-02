package main

import (
   "fmt"
   "log"
   "net/url"
)

func main() {
   s := "month=May&day=Friday"
   m, e := url.ParseQuery(s)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
