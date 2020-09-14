package main
import "fmt"

func main() {
   a := []int{10, 11}
   // example 1
   a = append(a, 12)
   // example 2
   a2 := []int{13, 14}
   a = append(a, a2...)
   // print
   fmt.Println(a)
}
