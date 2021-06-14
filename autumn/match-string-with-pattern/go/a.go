package main

import (
   "fmt"
   "regexp"
)

func main() {
   pattern, s := "o..", "north south"
   b, err := regexp.MatchString(pattern, s)
   fmt.Println(b, err)
}
