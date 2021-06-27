package main

import (
   "fmt"
   "strings"
)

func main() {
   a := strings.Fields(" north  south ")
   fmt.Printf("%q\n", a) // ["north" "south"]
}
