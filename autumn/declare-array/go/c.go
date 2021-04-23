package main
import "fmt"

func main() {
   { // example 1
      var s []int
      s = append(s, 10, 11)
      fmt.Println(s)
   }
   { // example 2
      var s = []int{10, 11}
      fmt.Println(s)
   }
   { // example 3
      s := []int{10, 11}
      fmt.Println(s)
   }
   { // example 4
      type slice []interface{}
      s := slice{
         slice{10, 11},
         slice{12, 13},
      }
      fmt.Println(s)
   }
}
