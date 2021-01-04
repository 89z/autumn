package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []map[string]int{
      {"month": 4, "day": 5},
      {"month": 5, "day": 4},
   }
   f := func (n, n2 int) bool {
      return a[n]["day"] < a[n2]["day"]
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
