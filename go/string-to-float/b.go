package main
import "fmt"

func main() {
   var n float64
   fmt.Sscanf("9.9", "%v", &n)
   fmt.Println(n == 9.9)
}
