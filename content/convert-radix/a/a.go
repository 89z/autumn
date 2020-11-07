package main
import "math/big"

func Decode(s string, n int) int64 {
   o := big.Int{}
   o.SetString(s, n)
   return o.Int64()
}

func main() {
   s := "q3ezbz"
   n := Decode(s, 36)
   println(n == 1577858399)
}
