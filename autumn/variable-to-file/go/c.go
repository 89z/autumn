package main

import (
   "os"
   "strings"
)

func main() {
   r_o := strings.NewReader("March\n")
   w_o, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   r_o.WriteTo(w_o)
}
