package main
import "encoding/ascii85"

func encode(src []byte) []byte {
   dst := make([]byte, ascii85.MaxEncodedLen(len(src)))
   return dst[:ascii85.Encode(dst, src)]
}

func main() {
   a := []byte{10, 11, 12, 13}
   b := encode(a)
   println(string(b) == "$4@7O")
}
