package main

import (
   "math/big"
   "time"
)

func idEncode(n int) string {
   t := time.Date(n, 1, 1, 0, 0, 0, 0, time.UTC)
   return big.NewInt(
      time.Since(t).Milliseconds(),
   ).Text(62)
}

func main() {
   s := idEncode(2021)
   println(s)
}
