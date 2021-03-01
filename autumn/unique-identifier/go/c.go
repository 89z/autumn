package main
import "strconv"

func encode(n int) string {
   return strconv.FormatInt(
      int64(n), 36,
   )
}

func main() {
   n := 0xFFFF_FFFF
   s := encode(n)
   println(s == "1z141z3")
}
