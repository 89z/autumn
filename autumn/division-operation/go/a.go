package main
import "math/bits"

func divMod(x, y uint32) (uint32, uint32) {
   return bits.Div32(0, x, y)
}

func main() {
   d, m := divMod(0xFFFF_FFFF, 2)
   println(d == 0x7FFF_FFFF, m == 1)
}
