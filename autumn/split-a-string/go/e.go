package main

import (
   "fmt"
   "strings"
)

func main() {
   s, sep := "west,north,east", ","
   a := strings.SplitN(s, sep, 2)
   fmt.Printf("%q\n", a) // ["west" "north,east"]
}
