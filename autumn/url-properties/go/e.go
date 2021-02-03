package main

import (
   "fmt"
   "net/url"
)

func main() {
   v := make(url.Values)
   v.Set("day", "Friday")
   fmt.Println(v)
}
