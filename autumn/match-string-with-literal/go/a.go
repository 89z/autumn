package main
import "bytes"

func main() {
   y := []byte("March")
   b := bytes.Contains(y, []byte("ar"))
   println(b)
}
