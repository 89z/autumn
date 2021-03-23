package main
import "fmt"

func main() {
   a := []byte{10, 11, 12, 13}
   // example 1
   s := fmt.Sprintf("%x", a)
   // example 2
   t := fmt.Sprintf("%X", a)
   // print
   println(s == "0a0b0c0d", t == "0A0B0C0D")
}
