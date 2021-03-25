package main
import "bytes"

func main() {
   b := []byte("March")
   c := bytes.Contains(b, []byte("ar"))
   println(c)
}
