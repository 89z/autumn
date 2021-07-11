package main

import (
   "fmt"
   "strings"
)

func main() {
   s := strings.Repeat("west east\n", 999)
   a := strings.SplitN(s, "\n", 2)
   fmt.Println(a)
}
