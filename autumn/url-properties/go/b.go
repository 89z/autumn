package main

import (
   "fmt"
   "net/url"
)

func main() {
   var u url.URL
   u.Path = "pkg"
   fmt.Printf("%+v\n", u)
}
