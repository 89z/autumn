package main
import "bytes"

func main() {
   b := []byte("north")
   c := bytes.Contains(b, []byte("or"))
   println(c)
}
