package main
import "fmt"

func main() {
   n := 9.9
   // example 1
   fmt.Printf("%.0f\n", n)
   // example 2
   s := fmt.Sprintf("%.0f", n)
   println(s)
}
