package main
import "fmt"

func main() {
   var n float64
   fmt.Sscan("9.9", &n)
   fmt.Println(n == 9.9)
}