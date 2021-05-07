package main

import (
   "fmt"
   "regexp"
)

func main() {
   b, e := regexp.MatchString("o..", "south north")
   fmt.Println(b, e)
}
