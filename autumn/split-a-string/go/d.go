package main

import (
   "fmt"
   "strings"
)

func main() {
   s := strings.Repeat("west east\n", 17)
   a := strings.FieldsFunc(s, func(r rune) bool {
      return r == '\n'
   })
   fmt.Println(a)
}
