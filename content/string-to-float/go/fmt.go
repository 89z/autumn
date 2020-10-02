package main
import "fmt"

func main() {
   // example 1
   n1 := float64(0)
   fmt.Sscanln("9.9", &n1)
   // example 2
   n2 := float64(0)
   fmt.Sscan("9.9", &n2)
   // example 3
   n3 := float64(0)
   fmt.Sscanf("9.9", "%v", &n3)
   // print
   println(n1 == 9.9, n2 == 9.9, n3 == 9.9)
}
