package main

import (
   "fmt"
   "regexp"
)

func main() {
   o := regexp.MustCompile("a.")
   // example 1
   s1 := o.FindString("January")
   // example 2
   a2 := o.FindAllString("January", -1)
   // print
   fmt.Println(s1, a2)
}
