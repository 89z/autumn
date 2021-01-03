package main

import (
   "log"
   "strconv"
)

func main() {
   s := "q3ezbz"
   n, e := strconv.ParseInt(s, 36, 0)
   if e != nil {
      log.Fatal(e)
   }
   println(n == 1609480799)
}
