package main
import "fmt"

func main() {
   { // example 1
      a := make([]int, 0)
      fmt.Println(a)
   }
   { // example 2
      a := make([]int, 0, 0)
      fmt.Println(a)
   }
}
