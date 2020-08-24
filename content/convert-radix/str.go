package main

import (
   "strings"
   "time"
)

type Radix64 struct {
   s_dig string
}

func NewRadix64() Radix64 {
   s := "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
   return Radix64{s}
}

func (o Radix64) Encode(n_in int) string {
   s_out := ""
   for {
      n_key := n_in % 64
      s_out = o.s_dig[n_key:n_key + 1] + s_out
      n_in /= 64
      if n_in == 0 {
         break
      }
   }
   return s_out
}

func (o Radix64) Decode(s_in string) int {
   n_out := 0
   for n_key := range s_in {
      s_chr := s_in[n_key:n_key + 1]
      n_out = n_out * 64 + strings.Index(o.s_dig, s_chr)
   }
   return n_out
}

func main() {
   n := int(time.Now().Unix())
   o := NewRadix64()
   s := o.Encode(n)
   n2 := o.Decode(s)
   println(n, s, n2 == n)
}
