package main
import "fmt"

func main() {
   { // example 1
      var a []int
      a = append(a, 10)
      fmt.Println(a)
   }
   { // example 2
      var a []int
      a = append(a, 10, 11)
      fmt.Println(a)
   }
   { // example 3
      a, b := []int{10, 11}, []int{12, 13}
      a = append(a, b...)
      fmt.Println(a)
   }
}
