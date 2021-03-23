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
   f := func (d, e int) bool {
      return a[d]["day"] < a[e]["day"]
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
