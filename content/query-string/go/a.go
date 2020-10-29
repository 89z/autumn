package main

import (
   "fmt"
   "net/url"
   "os"
)

func main() {
   m, e := url.ParseQuery("one=odd&two=even")
   if e != nil {
      os.Exit(1)
   }
   fmt.Println(m)
}
