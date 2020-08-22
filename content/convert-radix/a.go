package main
import (
   "math/big"
   "time"
)
// example 1
func r64_encode(n_in int64) string {
   o_big := big.NewInt(n_in)
   return o_big.Text(62)
}
// example 2
func r64_decode(s_in string) int64 {
   o_big := big.Int{}
   o_big.SetString(s_in, 62)
   return o_big.Int64()
}
// print
func main() {
   n1 := time.Now().Unix()
   s1 := r64_encode(n1)
   n2 := r64_decode(s1)
   println(n1, s1, n2 == n1)
}
