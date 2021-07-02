package main
import "encoding/hex"

func main() {
   b := []byte{10, 11, 12, 13}
   s := hex.EncodeToString(b)
   println(s == "0a0b0c0d")
}
