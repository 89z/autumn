package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []map[string]int{
      {"year": 2019, "month": 11},
      {"year": 2018, "month": 12},
   }
   f := func (n, n2 int) bool {
      return a[n]["year"] < a[n2]["year"]
   }
   sort.Slice(a, f)
   fmt.Println(a)
}
