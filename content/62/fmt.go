package main
import "fmt"

func main() {
   n, n2 := 0, 0
   // example 1
   fmt.Sscanln("10 11", &n, &n2)
   println(n == 10, n2 == 11)
   // example 2
   fmt.Sscan("10\n11", &n, &n2)
   println(n == 10, n2 == 11)
   // example 3
   fmt.Sscanf("10 11", "%v %v", &n, &n2)
   println(n == 10, n2 == 11)
}
