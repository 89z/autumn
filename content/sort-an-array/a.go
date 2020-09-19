package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []string{"BBBB", "AA", "CCC"}
   f := func (n, n1 int) bool {
      s, s1 := a[n], a[n1]
      return len(s) < len(s1)
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
