package main
import "fmt"

func main() {
   a := []byte{10, 11, 12}
   s := fmt.Sprintf("%x", a)
   println(s == "0a0b0c")
}
