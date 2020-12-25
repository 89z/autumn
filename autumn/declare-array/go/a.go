package main
import "fmt"

func main() {
   // example 1
   var a1 []int
   a1 = append(a1, 10, 11)
   // example 2
   var a2 = []int{10, 11}
   // example 3
   a3 := []int{10, 11}
   // example 4
   type A []interface{}
   a4 := A{
      A{10, 11},
      A{12, 13},
   }
   // print
   fmt.Println(a1, a2, a3, a4)
}
