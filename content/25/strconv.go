package main

import (
   "os"
   "strconv"
)

func main() {
   s := "2.9"
   n, e := strconv.ParseFloat(s, 64)
   if e != nil {
      os.Exit(1)
   }
   println(n == 2.9)
}
