package main

import (
   "os"
   "strconv"
)

func main() {
   s := "9.9"
   n, e := strconv.ParseFloat(s, 64)
   if e != nil {
      os.Exit(1)
   }
   println(n == 9.9)
}
