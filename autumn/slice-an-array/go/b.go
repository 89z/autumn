package main
import "fmt"

func main() {
   a := []int{0, 10, 20, 30, 40}
   { // example 1
      n := a[1]
      fmt.Println(n)
   }
   { // example 2
      b := a[2:]
      fmt.Println(b)
   }
   { // example 3
      b := a[2:4]
      fmt.Println(b)
   }
}
