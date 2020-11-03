package main

import (
   "fmt"
   "math/big"
)

func main() {
   nX, nY := big.NewInt(7), big.NewInt(19)
   nX.Exp(nX, nY, nil)
   fmt.Println(nX)
}
