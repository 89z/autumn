package main

import (
   "math/big"
   "time"
   "fmt"
)

func IdDecode(s string, year int) time.Time {
   z := big.Int{}
   z.SetString(s, 36)
   return time.Date(year, 1, 1, 0, 0, int(z.Int64()), 0, time.UTC)
}

func main() {
   o := IdDecode("itrzz", 2020)
   fmt.Println(o)
}
