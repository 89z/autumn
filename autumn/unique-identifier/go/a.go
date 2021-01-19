package main

import (
   "fmt"
   "math/big"
   "time"
)

func IdDecode(s string, year int) time.Time {
   x := new(big.Int)
   x.SetString(s, 36)
   d := time.Duration(x.Int64()) * time.Second
   return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Add(d)
}

func main() {
   t := IdDecode("6dv3d", 2020)
   fmt.Println(t)
}
