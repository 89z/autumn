package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "west,north,east"
   a := strings.Split(s, ",")
   fmt.Printf("%q\n", a)
}
