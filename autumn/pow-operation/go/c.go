package main

import (
   "fmt"
   "math/big"
)

func pow(a, b int64) *big.Int {
   c, d := big.NewInt(a), big.NewInt(b)
   return new(big.Int).Exp(c, d, nil)
}

func main() {
   n := pow(10, 5)
   fmt.Println(n) // 100000
}
