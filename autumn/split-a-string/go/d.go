package main

import (
   "fmt"
   "strings"
)

func comma(r rune) bool {
   return r == ','
}

func main() {
   a := strings.FieldsFunc(",north,,south,", comma)
   fmt.Printf("%q\n", a) // ["north" "south"]
}
