package main

import (
   "io"
   "os"
   "strings"
)

func main() {
   in_o, e := os.Open("a.go")
   if e != nil {
      os.Exit(1)
   }
   out_o := strings.Builder{}
   io.Copy(&out_o, in_o)
   s := out_o.String()
   print(s)
}
