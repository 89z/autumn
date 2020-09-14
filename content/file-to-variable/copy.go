package main

import (
   "io"
   "os"
   "strings"
)

func main() {
   o, e := os.Open("a.txt")
   if e != nil {
      os.Exit(1)
   }
   o2 := strings.Builder{}
   io.Copy(&o2, o)
   s := o2.String()
   print(s)
}
