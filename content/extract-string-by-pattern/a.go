package main

import (
   "fmt"
   "regexp"
)

func main() {
   o := regexp.MustCompile("a.")
   // example 1
   s := o.FindString("January")
   // example 2
   a := o.FindAllString("January", -1)
   // print
   fmt.Println(s, a)
}
