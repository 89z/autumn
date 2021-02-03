package main

import (
   "fmt"
   "net/url"
)

func main() {
   var u url.URL
   u.Host = "golang.org"
   fmt.Printf("%+v\n", u)
}
