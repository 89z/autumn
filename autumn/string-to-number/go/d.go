package main

import (
   "log"
   "strconv"
)

func main() {
   n, e := strconv.ParseFloat("99.9", 64)
   if e != nil {
      log.Fatal(e)
   }
   println(n == 99.9)
}
