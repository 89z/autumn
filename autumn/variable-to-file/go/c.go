package main

import (
   "os"
   "strings"
)

func main() {
   in_o := strings.NewReader("March\n")
   out_o, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   in_o.WriteTo(out_o)
}
