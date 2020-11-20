package main
import "fmt"

func main() {
   var n int
   fmt.Sscan("100", &n)
   fmt.Println(n == 100)
}
