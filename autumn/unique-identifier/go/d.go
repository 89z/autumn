package main

import (
   "log"
   "strconv"
)

func main() {
   n, e := strconv.ParseInt("qm8rbz", 36, 0)
   if e != nil {
      log.Fatal(e)
   }
   println(n == 1609480799)
}
