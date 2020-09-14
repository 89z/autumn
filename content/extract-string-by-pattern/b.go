package main

import (
   "fmt"
   "regexp"
)

func main() {
   o := regexp.MustCompile("a(..)")
   // example 1
   a := o.FindStringSubmatch("January")
   // example 2
   a2 := o.FindAllStringSubmatch("January", -1)
   // print
   fmt.Println(a, a2)
}
