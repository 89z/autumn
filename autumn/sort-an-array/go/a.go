package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []map[string]int{
      {"month": 10, "day": 31}, {"month": 11, "day": 30},
   }
   f := func (d, e int) bool {
      return a[d]["day"] < a[e]["day"]
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
