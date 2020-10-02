package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "May June"
   // example 1
   a1 := strings.Split(s, " ")
   // example 2
   a2 := strings.Fields(s)
   // print
   fmt.Println(a1, a2)
}
