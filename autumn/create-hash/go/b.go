package main

import (
   "crypto/md5"
   "fmt"
)

func sum(s string) string {
   b := md5.Sum([]byte(s))
   return fmt.Sprintf("%x", b)
}

func main() {
   s := sum("south north")
   fmt.Println(s)
}
