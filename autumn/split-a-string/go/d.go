package main

import (
   "fmt"
   "strings"
)

func main() {
   a := strings.Split(",north,,south,", ",")
   fmt.Printf("%q\n", a) // ["" "north" "" "south" ""]
}
