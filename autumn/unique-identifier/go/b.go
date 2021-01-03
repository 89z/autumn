package main
import "math/big"

func Encode36(x int) string {
   n := int64(x)
   return big.NewInt(n).Text(36)
}

func main() {
   n := 1609480799
   s := Encode36(n)
   println(s == "qm8rbz")
}
