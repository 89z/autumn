package main

import (
   "strings"
   "time"
)

type Radix struct {
   Digits string
   Size int64
}

func NewRadix64() Radix {
   s := "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
   n := len(s)
   return Radix { s, int64(n) }
}

func (o Radix) Encode(n_in int64) string {
   s_out := ""
   for {
      n_key := n_in % o.Size
      s_out = o.Digits[n_key:n_key + 1] + s_out
      n_in /= o.Size
      if n_in == 0 {
         break
      }
   }
   return s_out
}

func (o Radix) Decode(s_in string) int64 {
   n_out := int64(0)
   for n_key := range s_in {
      s_chr := s_in[n_key:n_key + 1]
      n_ind := strings.Index(o.Digits, s_chr)
      n_out = n_out * o.Size + int64(n_ind)
   }
   return n_out
}

func main() {
   n := time.Now().Unix()
   o := NewRadix64()
   s := o.Encode(n)
   n2 := o.Decode(s)
   println(n, s, n2 == n)
}
