package main

import (
   "crypto/md5"
   "encoding/hex"
   "io"
)

func sum(s string) string {
   h := md5.New()
   io.WriteString(h, s)
   return hex.EncodeToString(h.Sum(nil))
}

func main() {
   s := sum("south north")
   println(s == "f53b1396fe2736a7e5c44629ee1a3af5")
}
