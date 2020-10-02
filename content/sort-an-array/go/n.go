package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []int{10, 9}
   sort.Ints(a)
   fmt.Println(a)
}
