package main
import "fmt"

func main() {
   { // example 1
      s := make([]int, 0)
      fmt.Println(s)
   }
   { // example 2
      s := make([]int, 0, 0)
      fmt.Println(s)
   }
}
