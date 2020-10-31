package main

import (
   "os"
   "strconv"
)

func main() {
   s := "q3ezbz"
   n, e := strconv.ParseInt(s, 36, 0)
   if e != nil {
      os.Exit(1)
   }
   println(n == 1577858399)
}
