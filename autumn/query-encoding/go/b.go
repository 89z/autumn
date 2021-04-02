package main

import (
   "fmt"
   "net/url"
)

func decode(s string) url.Values {
   u := url.URL{RawQuery: s}
   return u.Query()
}

func main() {
   m := decode("west=left&east=right")
   fmt.Println(m)
}
