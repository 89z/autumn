package main

import (
   "fmt"
   "regexp"
)

func main() {
   o := regexp.MustCompile("e(..)")
   // example 1
   a := o.FindStringSubmatch("Wednesday")
   // example 2
   a2 := o.FindAllStringSubmatch("Wednesday", -1)
   // print
   fmt.Println(a, a2)
}
