package main
import "fmt"

func main() {
   n, n2 := 0, 0
   // example 1
   fmt.Sscanln("8 9", &n, &n2)
   println(n == 8, n2 == 9)
   // example 2
   fmt.Sscan("8\n9", &n, &n2)
   println(n == 8, n2 == 9)
   // example 3
   fmt.Sscanf("8 9", "%v %v", &n, &n2)
   println(n == 8, n2 == 9)
}
