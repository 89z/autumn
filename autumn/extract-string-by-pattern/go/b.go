package main

import (
   "fmt"
   "regexp"
)

func main() {
   a := regexp.MustCompile("a(..)").FindAllStringSubmatch("January", -1)
   fmt.Println(a)
}
