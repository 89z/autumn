package main

import (
   "fmt"
   "strings"
)

func main() {
   a := strings.Split("west,north,east", ",")
   fmt.Println(a) // [west north east]
}
