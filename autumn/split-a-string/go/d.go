package main

import (
   "fmt"
   "strings"
)

func main() {
   a := strings.SplitN("west,north,east", ",", 2)
   fmt.Println(a) // [west north,east]
}
