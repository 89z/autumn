package main
import "fmt"

func main() {
   a := []int{10, 11}
   // example 1
   a1 := append(a, 12)
   // example 2
   a2a := []int{12, 13}
   a2 := append(a, a2a...)
   // print
   fmt.Println(a1, a2)
}
