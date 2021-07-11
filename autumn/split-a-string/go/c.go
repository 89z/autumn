package main

import (
   "fmt"
   "strings"
)

func main() {
   s := strings.Repeat("west east\n", 24)
   a := strings.Fields(s)
   fmt.Println(a)
}
