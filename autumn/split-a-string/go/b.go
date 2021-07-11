package main

import (
   "fmt"
   "strings"
)

func main() {
   s := strings.Repeat("west east\n", 48)
   a := strings.Split(s, "\n")
   fmt.Println(a)
}
