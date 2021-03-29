package main

import (
   "fmt"
   "log"
   "net/url"
)

func main() {
   s := "west=left&east=right"
   m, e := url.ParseQuery(s)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
