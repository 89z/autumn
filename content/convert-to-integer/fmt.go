package main
import "fmt"
func main() {
   n1 := 0
   n2 := 0
   // example 1
   fmt.Sscanln("10 11", &n1, &n2)
   println(n1, n2)
   // example 2
   fmt.Sscan("10\n11", &n1, &n2)
   println(n1, n2)
   // example 3
   fmt.Sscanf("10 11", "%v %v", &n1, &n2)
   println(n1, n2)
}
