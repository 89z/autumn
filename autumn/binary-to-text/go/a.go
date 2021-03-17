package main
import "fmt"
import "encoding/hex"

func main() {
   s := fmt.Sprintf("%x", []byte{10, 10})
   println(s == "0a0a")
   println(hex.EncodeToString([]byte{10,10}))
}
