package main
import "fmt"

func main() {
   s := []int{0, 10, 20, 30, 40}
   { // example 1
      n := s[1]
      fmt.Println(n)
   }
   { // example 2
      b := s[2:]
      fmt.Println(b)
   }
   { // example 3
      b := s[2:4]
      fmt.Println(b)
   }
}
