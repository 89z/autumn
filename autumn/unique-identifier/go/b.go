package main

import (
   "math/big"
   "time"
)

func ID(year int, d time.Duration) string {
   time.Sleep(d)
   t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
   x := time.Since(t).Nanoseconds() / int64(d)
   return big.NewInt(x).Text(62)
}

func main() {
   s := ID(2021, time.Millisecond)
   println(s)
}
