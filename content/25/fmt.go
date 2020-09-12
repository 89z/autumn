package main
import "fmt"

func main() {
   n, n2 := float64(0), float64(0)
   // example 1
   fmt.Sscanln("9.9 9.8", &n, &n2)
   println(n == 9.9, n2 == 9.8)
   // example 2
   fmt.Sscan("9.9\n9.8", &n, &n2)
   println(n == 9.9, n2 == 9.8)
   // example 3
   fmt.Sscanf("9.9 9.8", "%v %v", &n, &n2)
   println(n == 9.9, n2 == 9.8)
}
