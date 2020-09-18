package main
import "strings"

type Radix64 struct {
   s_dig string
}

func NewRadix64() Radix64 {
   s := "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
   return Radix64{s}
}

func (o_rad Radix64) Encode(n_in int) string {
   s_out := ""
   for {
      n_key := n_in % 64
      s_out = o_rad.s_dig[n_key:n_key + 1] + s_out
      n_in /= 64
      if n_in == 0 {
         break
      }
   }
   return s_out
}

func (o_rad Radix64) Decode(s_in string) int {
   n_out := 0
   for n_key := range s_in {
      s_chr := s_in[n_key:n_key + 1]
      n_out = n_out * 64 + strings.Index(o_rad.s_dig, s_chr)
   }
   return n_out
}

func main() {
   n := 1577858399
   o := NewRadix64()
   // example 1
   s1 := o.Encode(n)
   n1 := o.Decode(s1)
   // example 2
   s2 := o.Encode(n - 1)
   n2 := o.Decode(s2)
   // print
   println(s1 == "0T22KU", n1 == n, s2 == "0T22KT", n2 == n - 1)
}
