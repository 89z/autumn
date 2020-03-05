package main
import (
   "fmt"
   "sort"
)
func main() {
   // example 1
   a1 := []int{20, 10}
   a2 := sort.IntSlice(a1)
   // example 2
   a3 := []string{"monday", "tue"}
   sort.Slice(a3, func(n1, n2 int) bool {
      return len(a3[n1]) < len(a3[n2])
   })
   // print
   fmt.Println(a2, a3)
}
