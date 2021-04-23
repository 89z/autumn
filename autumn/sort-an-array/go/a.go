package main

import (
   "fmt"
   "sort"
)

func main() {
   s := []map[string]int{
      {"month": 10, "day": 31}, {"month": 11, "day": 30},
   }
   f := func (d, e int) bool {
      return s[d]["day"] < s[e]["day"]
   }
   sort.Slice(s, f)
   fmt.Println(s)
}
