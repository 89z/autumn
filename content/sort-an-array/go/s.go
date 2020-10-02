package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []string{"April", "May", "June"}
   // example 1
   sort.Strings(a)
   fmt.Println(a)
   // example 2
   f2 := func (n, n2 int) bool {
      s, s2 := a[n], a[n2]
      return len(s) < len(s2)
   }
   sort.Slice(a, f2)
   fmt.Println(a)
}
