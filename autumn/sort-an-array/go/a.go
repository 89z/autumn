package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []map[]int{
      {"month": 11, "day": 30},
      {"month": 10, "day": 31},
   }
   f := func (n, n2 int) bool {
      return a[n]["month"] < a[n2]["month"]
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
