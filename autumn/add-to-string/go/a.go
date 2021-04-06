package main

import (
   "fmt"
   "strings"
)

func main() {
   b := new(strings.Builder)
   fmt.Fprintln(b, "south north")
   print(b.String())
}
