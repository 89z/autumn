package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "May,June,July"
   a := strings.Split(s, ",")
   fmt.Printf("%q\n", a)
}
