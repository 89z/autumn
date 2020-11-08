package main

import (
   "fmt"
   "net/url"
   "os"
)

func main() {
   o, e := url.Parse("https://example.com/one?two=even")
   if e != nil {
      os.Exit(1)
   }
   m := o.Query()
   fmt.Println(m)
}
