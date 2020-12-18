package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "May,June,July"
   a := strings.SplitN(s, ",", 2)
   fmt.Printf("%q\n", a)
}
