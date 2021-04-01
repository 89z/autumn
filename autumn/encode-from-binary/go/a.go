package main
import "encoding/base64"

func main() {
   b := []byte{10, 11, 12, 13}
   s := base64.StdEncoding.EncodeToString(b)
   println(s == "CgsMDQ==")
}
