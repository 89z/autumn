package main

import (
   "math/big"
   "time"
)

func IdEncode(year int) string {
   t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
   x := time.Since(t).Milliseconds()
   return big.NewInt(x).Text(62)
}

func main() {
   s := IdEncode(2021)
   println(s)
}
