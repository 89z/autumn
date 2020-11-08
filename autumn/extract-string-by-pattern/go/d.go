package main

import (
   "fmt"
   "regexp"
)

func main() {
   a := regexp.MustCompile("a(..)").FindStringSubmatch("January")
   fmt.Println(a)
}
