package main

import (
   "fmt"
   "net/url"
)

func main() {
   s := "west=left&east=right"
   m, e := url.ParseQuery(s)
   if e != nil {
      panic(e)
   }
   fmt.Println(m)
}
