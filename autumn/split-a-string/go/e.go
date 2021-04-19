package main

import (
   "fmt"
   "strings"
)

func main() {
   s := "west,north,east"
   a := strings.SplitN(s, ",", 2)
   fmt.Printf("%q\n", a)
}
