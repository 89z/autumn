package main

import (
   "log"
   "strconv"
)

func main() {
   n, e := strconv.Atoi("100")
   if e != nil {
      log.Fatal(e)
   }
   println(n == 100)
}
