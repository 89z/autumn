package main

import (
   "math/big"
   "time"
)

type Radix struct {
   Size int
}

func (o Radix) Encode(n_in int64) string {
   o_big := big.NewInt(n_in)
   return o_big.Text(o.Size)
}

func (o Radix) Decode(s_in string) int64 {
   o_big := big.Int{}
   o_big.SetString(s_in, o.Size)
   return o_big.Int64()
}

func main() {
   n := time.Now().Unix()
   o := Radix{62}
   s := o.Encode(n)
   n2 := o.Decode(s)
   println(n, s, n2 == n)
}
