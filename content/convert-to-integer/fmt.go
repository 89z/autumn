package main
import "fmt"

func main() {
   // example 1
   n1 := 0
   fmt.Sscanln("10", &n1)
   // example 2
   n2 := 0
   fmt.Sscan("10", &n2)
   // example 3
   n3 := 0
   fmt.Sscanf("10", "%v", &n3)
   // print
   println(n1 == 10, n2 == 10, n3 == 10)
}
