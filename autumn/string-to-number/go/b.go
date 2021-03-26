package main
import "fmt"

func main() {
   { // int
      var n int
      fmt.Sscanf("100", "%v", &n)
      fmt.Println(n == 100)
   }
   { // float
      var n float64
      fmt.Sscanf("99.9", "%v", &n)
      fmt.Println(n == 99.9)
   }
}
