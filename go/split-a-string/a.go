package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "May June"
   a := strings.Fields(s)
   fmt.Printf("%q\n", a)
}
