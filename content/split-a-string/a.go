package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "May June"
   // example A
   aA := strings.Split(s, " ")
   // example B
   aB := strings.Fields(s)
   // print
   fmt.Println(aA, aB)
}
