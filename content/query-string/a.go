package main

import (
   "fmt"
   "net/url"
   "os"
)

func main() {
   // example 1
   m, e := url.ParseQuery("one=odd&two=even")
   if e != nil {
      os.Exit(1)
   }
   // example 2
   m2 := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   s := m2.Encode()
   // print
   fmt.Println(m, s)
}
