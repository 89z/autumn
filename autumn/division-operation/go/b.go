package main
import "math/bits"

func divMod(x, y uint64) (uint64, uint64) {
   return bits.Div64(0, x, y)
}

func main() {
   d, m := divMod(0xFFFF_FFFF_FFFF_FFFF, 2)
   println(d == 0x7FFF_FFFF_FFFF_FFFF, m == 1)
}
