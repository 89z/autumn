package main
import "encoding/base64"

func main() {
   a := []byte{10, 11, 12}
   s := base64.StdEncoding.EncodeToString(a)
   println(s == "CgsM")
}
