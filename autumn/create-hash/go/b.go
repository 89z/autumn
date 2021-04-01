package main

import (
   "crypto/md5"
   "encoding/hex"
)

func sum(s string) string {
   b := md5.Sum([]byte(s))
   return hex.EncodeToString(b[:])
}

func main() {
   s := sum("south north")
   println(s == "f53b1396fe2736a7e5c44629ee1a3af5")
}
