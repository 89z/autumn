package main

import (
   "fmt"
   "math/big"
)

func pow(x, y int64) *big.Int {
   return new(big.Int).Exp(
      big.NewInt(x), big.NewInt(y), nil,
   )
}

func main() {
   n := pow(10, 5)
   fmt.Println(n)
}
