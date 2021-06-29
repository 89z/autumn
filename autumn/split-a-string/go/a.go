package main

import (
   "fmt"
   "strings"
)

func main() {
   a := strings.SplitN(",north,,south,", ",", 2)
   fmt.Printf("%q\n", a) // ["" "north,,south,"]
}
