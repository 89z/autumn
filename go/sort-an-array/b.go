package main

import (
   "fmt"
   "sort"
)

func main() {
   a := []string{"May", "June"}
   sort.Strings(a)
   fmt.Println(a)
}
