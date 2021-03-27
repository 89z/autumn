package main
import "fmt"

func pop(a *[]int) int {
   d := len(*a)
   e := (*a)[d - 1]
   *a = (*a)[:d - 1]
   return e
}

func main() {
   a := []int{0, 10, 20}
   n := pop(&a)
   fmt.Println(a, n)
}
