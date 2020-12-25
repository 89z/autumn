package main
import "fmt"

func main() {
   // example 1
   var a1 = make([]int, 0)
   a1 = append(a1, 10, 11)
   // example 2
   a2 := make([]int, 0)
   a2 = append(a2, 10, 11)
   // print
   fmt.Println(a1, a2)
}
