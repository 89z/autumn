package main
import "fmt"

func main() {
   var n int
   fmt.Sscanf("100", "%v", &n)
   fmt.Println(n == 100)
}
