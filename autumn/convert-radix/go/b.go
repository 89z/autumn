package main
import "math/big"

func Encode36(x int) string {
   n := int64(x)
   return big.NewInt(n).Text(36)
}

func main() {
   n := 1577858399
   s := Encode36(n)
   println(s == "q3ezbz")
}
