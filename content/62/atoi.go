package main

import (
   "os"
   "strconv"
)

func main() {
   n, e := strconv.Atoi("10")
   if e != nil {
      os.Exit(1)
   }
   println(n == 10)
}
