package main

import (
   "fmt"
   "net/url"
)

func main() {
   v := make(url.Values)
   v.Set("west", "left")
   fmt.Println(v)
}
