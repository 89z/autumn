package main

import (
   "fmt"
   "regexp"
)

func main() {
   o := regexp.MustCompile("e.")
   // example 1
   s := o.FindString("Wednesday")
   // example 2
   a := o.FindAllString("Wednesday", -1)
   // print
   fmt.Println(s, a)
}
