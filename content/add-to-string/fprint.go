package main

import (
   "fmt"
   "strings"
)

func main() {
   o := strings.Builder{}
   fmt.Fprint(&o, "May")
   s := o.String()
   println(s == "May")
}
