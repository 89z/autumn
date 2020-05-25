package main
import (
   "strings"
   "time"
)
var s_dig = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
// example 1
func r64_encode(n_in int) string {
   s_out := ""
   for n_in > 0 {
      n_key := n_in % 64
      s_out = s_dig[n_key:n_key + 1] + s_out
      n_in /= 64
   }
   return s_out
}
// example 1
func r64_decode(s_in string) int {
   n_out := 0
   for n_key := range s_in {
      s_chr := s_in[n_key:n_key + 1]
      n_out = n_out * 64 + strings.Index(s_dig, s_chr)
   }
   return n_out
}
// print
func main() {
   n1 := int(time.Now().Unix())
   s1 := r64_encode(n1)
   n2 := r64_decode(s1)
   println(n1, s1, n2 == n1)
}
