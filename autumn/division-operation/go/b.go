package main
import "math/bits"

func main() {
   n := uint64(18446744073709551615)
   q, r := bits.Div64(0, n, 2)
   println(q == n / 2, r == 1)
}
