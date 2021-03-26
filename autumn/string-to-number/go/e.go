package main

import (
   "log"
   "strconv"
)

func main() {
   n, e := strconv.ParseInt("100", 10, 64)
   if e != nil {
      log.Fatal(e)
   }
   println(n == 100)
}
