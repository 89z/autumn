package main

import (
   "io"
   "strings"
)

func main() {
   b := new(strings.Builder)
   io.WriteString(b, "north")
   s := b.String()
   println(s == "north")
}
