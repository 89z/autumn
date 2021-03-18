package main
import "fmt"

func main() {
   s := fmt.Sprintf("%x", []byte{10, 10})
   println(s == "0a0a")
}
