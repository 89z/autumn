package main
import "fmt"

func prepend(a []int, n int) []int {
   a = append(a, 0)
   copy(a[1:], a)
   a[0] = n
   return a
}

func main() {
   a := []int{10, 11}
   a = prepend(a, 9)
   fmt.Println(a)
}
