package main

import (
   "fmt"
   "math/big"
)

func pow(d, e int64) *big.Int {
   return new(big.Int).Exp(
      big.NewInt(d), big.NewInt(e), nil,
   )
}

func main() {
   n := pow(10, 5)
   fmt.Println(n)
}
