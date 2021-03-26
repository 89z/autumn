package main
import "fmt"

func main() {
   { // int
      var n int
      fmt.Sscan("100", &n)
      fmt.Println(n == 100)
   }
   { // float
      var n float64
      fmt.Sscan("99.9", &n)
      fmt.Println(n == 99.9)
   }
}
