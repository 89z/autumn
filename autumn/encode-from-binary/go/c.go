package main
import "fmt"

func main() {
   b := []byte{10, 11, 12, 13}
   { // example 1
      s := fmt.Sprintf("%x", b)
      fmt.Println(s == "0a0b0c0d")
   }
   { // example 2
      s := fmt.Sprintf("%X", b)
      fmt.Println(s == "0A0B0C0D")
   }
}
