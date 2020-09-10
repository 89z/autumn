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
   o2 := o.Query()
   fmt.Println(o2)
}
