package main

import (
   "fmt"
   "regexp"
)

func main() {
   a := regexp.MustCompile("a.").FindAllString("January", -1)
   fmt.Println(a)
}
