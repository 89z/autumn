package main

import (
   "fmt"
   "strconv"
)

func main() {
   b := strconv.AppendInt(nil, 0xFF, 10)
   fmt.Printf("%c\n", b) // [2 5 5]
}
