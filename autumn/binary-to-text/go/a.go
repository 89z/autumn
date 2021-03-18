package main
import "encoding/hex"

func main() {
   println(hex.EncodeToString([]byte{10,10}))
}
