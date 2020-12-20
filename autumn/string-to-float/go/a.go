package main
import "fmt"

func FloatVal(s string) float64 {
   var n float64
   fmt.Sscan(s, &n)
   return n
}

func main() {
   n := FloatVal("9.9")
   fmt.Println(n == 9.9)
}
