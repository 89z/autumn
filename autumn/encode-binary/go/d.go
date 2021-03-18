package main
import "encoding/hex"

func main() {
   a := []byte{10, 11, 12}
   s := hex.EncodeToString(a)
   println(s == "0a0b0c")
}
