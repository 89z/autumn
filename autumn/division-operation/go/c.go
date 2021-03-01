package main
import "math/bits"

func divMod(x, y uint) (uint, uint) {
   return bits.Div(0, x, y)
}

func main() {
   d, m := divMod(0xFFFF_FFFF, 2)
   println(d == 0x7FFF_FFFF, m == 1)
}
