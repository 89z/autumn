package main
import "fmt"

func main() {
   { // example 1
      var a []int
      a = append(a, 10, 11)
      fmt.Println(a)
   }
   { // example 2
      var a = []int{10, 11}
      fmt.Println(a)
   }
   { // example 3
      a := []int{10, 11}
      fmt.Println(a)
   }
   { // example 4
      type slice []interface{}
      a := slice{
         slice{10, 11},
         slice{12, 13},
      }
      fmt.Println(a)
   }
}
