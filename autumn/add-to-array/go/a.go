package main
import "fmt"

func main() {
   var a []int
   // example 1
   a = append(a, 10)
   // example 2
   a = append(a, 11, 12)
   // example 3
   a3 := []int{13, 14}
   a = append(a, a3...)
   // print
   fmt.Println(a)
}
