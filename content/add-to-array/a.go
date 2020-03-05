package main
import "fmt"
func main() {
   // example 1
   var a1 = []int{10, 20}
   a1 = append(a1, 30)
   // example 2
   var a2 = []int{10, 20}
   a2 = append(a2, []int{30, 40}...)
   // print
   fmt.Println(a1, a2)
}
