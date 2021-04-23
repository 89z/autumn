package main
import "fmt"

func pop(s *[]int) int {
   d := len(*s)
   e := (*s)[d - 1]
   *s = (*s)[:d - 1]
   return e
}

func main() {
   s := []int{0, 10, 20}
   n := pop(&s)
   fmt.Println(s, n)
}
