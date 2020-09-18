package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []string{"BBBB", "AA", "CCC"}
   f := func (nA, nB int) bool {
      sA, sB := a[nA], a[nB]
      return len(sA) < len(sB)
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
