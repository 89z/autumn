package main
import "math/bits"

func main() {
   n := uint32(4294967295)
   q, r := bits.Div32(0, n, 2)
   println(q == n / 2, r == 1)
}
