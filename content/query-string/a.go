package main

import (
   "fmt"
   "net/url"
   "os"
)

func main() {
   // example A
   m, e := url.ParseQuery("one=odd&two=even")
   if e != nil {
      os.Exit(1)
   }
   // example B
   mB := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   s := mB.Encode()
   // print
   fmt.Println(m, s)
}
