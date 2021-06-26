package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []string{"west", "east"}
   sort.Strings(a)
   fmt.Println(a)
}
