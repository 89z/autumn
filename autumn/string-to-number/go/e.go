package main

import (
   "log"
   "strconv"
)

func main() {
   { // int
      n, e := strconv.ParseInt("100", 10, 0)
      if e != nil {
         log.Fatal(e)
      }
      println(n == 100)
   }
   { // int64
      n, e := strconv.ParseInt("100", 10, 64)
      if e != nil {
         log.Fatal(e)
      }
      println(n == 100)
   }
}
