package main

import (
   "fmt"
   "math/big"
)

func main() {
   n := big.NewInt(0xFF)
   fmt.Println(n.Bytes()) // [255]
}
