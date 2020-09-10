package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []string{"BBBB", "AA", "CCC"}
   f := func (n, n2 int) bool {
      s, s2 := a[n], a[n2]
      return len(s) < len(s2)
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
