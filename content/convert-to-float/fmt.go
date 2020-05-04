package main
import "fmt"
func main() {
   var n1, n2 float64
   // example 1
   fmt.Sscanln("1.8 1.9", &n1, &n2)
   println(n1, n2)
   // example 2
   fmt.Sscan("1.8\n1.9", &n1, &n2)
   println(n1, n2)
   // example 3
   fmt.Sscanf("1.8 1.9", "%v %v", &n1, &n2)
   println(n1, n2)
}
