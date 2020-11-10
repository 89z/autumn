package main

import (
   "os"
   "strings"
)

func main() {
   oR := strings.NewReader("March\n")
   oW, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   oR.WriteTo(oW)
}
