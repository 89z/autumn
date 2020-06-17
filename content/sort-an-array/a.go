package main
import (
   "fmt"
   "sort"
)
func main() {
   a := []string{"BBBB", "AA", "CCC"}
   f := func(n1, n2 int) bool {
      return len(a[n1]) < len(a[n2])
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
