package main
import "fmt"

func main() {
   var a []int
   // example 1
   a = append(a, 10)
   // example 2
   a2 := []int{11, 12}
   a = append(a, a2...)
   // print
   fmt.Println(a)
}
