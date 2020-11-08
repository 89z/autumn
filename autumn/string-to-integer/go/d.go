package main

import (
   "os"
   "strconv"
)

func main() {
   s := "11"
   n, e := strconv.ParseInt(s, 10, 0)
   if e != nil {
      os.Exit(1)
   }
   println(n == 11)
}
