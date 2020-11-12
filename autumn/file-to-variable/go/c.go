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
   o1 := strings.Builder{}
   io.Copy(&o1, o)
   s := o1.String()
   print(s)
}
