package main

import (
   "os"
   "strings"
)

func split(s string) (string, string) {
   f := strings.LastIndexByte
   d, e := f(s, os.PathSeparator), f(s, '.')
   return s[d + 1:e], s[e + 1:]
}

func main() {
   stem, ext := split(`C:\go\README.md`)
   println(stem == "README", ext == "md")
}
