package main

import (
   "fmt"
   "math/big"
)

func main() {
   n := new(big.Int)
   n.Exp(big.NewInt(10), big.NewInt(5), nil)
   fmt.Println(n)
}
