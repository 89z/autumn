package main
import "fmt"

func main() {
   n, n2 := float64(0), float64(0)
   // example 1
   fmt.Sscanln("2.8 2.9", &n, &n2)
   println(n == 2.8, n2 == 2.9)
   // example 2
   fmt.Sscan("2.8\n2.9", &n, &n2)
   println(n == 2.8, n2 == 2.9)
   // example 3
   fmt.Sscanf("2.8 2.9", "%v %v", &n, &n2)
   println(n == 2.8, n2 == 2.9)
}
