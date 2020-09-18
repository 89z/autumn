package main

import (
   "math/big"
   "time"
)

type Radix struct {
   n_len int
}

func (o Radix) Encode(n_in int64) string {
   o_big := big.NewInt(n_in)
   return o_big.Text(o.n_len)
}

func (o Radix) Decode(s_in string) int64 {
   o_big := big.Int{}
   o_big.SetString(s_in, o.n_len)
   return o_big.Int64()
}

func main() {
   o := Radix{62}
   n := time.Now().Unix()
   // example 1
   s1 := o.Encode(n)
   // example 2
   n2 := o.Decode(s1)
   // print
   println(n, s1, n2 == n)
}
