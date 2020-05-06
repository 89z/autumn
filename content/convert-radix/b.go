package main
import (
   "strings"
   "time"
)
var s_dig = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
// example 1
func r64_encode(n_in int) string {
   var s_out string
   for n_in > 0 {
      n_key := n_in & 63
      s_out = s_dig[n_key:n_key + 1] + s_out
      n_in >>= 6
   }
   return s_out
}
// example 1
func r64_decode(s_in string) int {
   var n_out int
   for _, s_chr := range s_in {
      n_out = n_out << 6 | strings.IndexRune(s_dig, s_chr)
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
