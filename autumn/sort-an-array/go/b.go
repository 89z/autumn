package main

import (
   "fmt"
   "sort"
)

func main() {
   s := []struct{month, day int}{
      {10, 31}, {11, 30},
   }
   f := func(d, e int) bool {
      return s[d].day < s[e].day
   }
   sort.SliceStable(s, f)
   fmt.Println(s)
}
