package main

import (
   "log"
   "strconv"
)

func main() {
   n, e := strconv.ParseInt("1z141z3", 36, 0)
   if e != nil {
      log.Fatal(e)
   }
   println(n == 0xFFFF_FFFF)
}
