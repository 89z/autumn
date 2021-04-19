package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "west east"
   a := strings.Fields(s)
   fmt.Printf("%q\n", a)
}
