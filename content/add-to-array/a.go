package main
import "fmt"
func main() {
   // example 1
   a1 := []int{10, 11}
   a1 = append(a1, 12)
   // example 2
   a2 := []int{10, 11}
   a2 = append(a2, []int{12, 13}...)
   // print
   fmt.Println(a1, a2)
}
