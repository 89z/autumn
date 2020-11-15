package main

import (
   "io"
   "os"
   "strings"
)

func GetContents(s string) (string, error) {
   open_o, e := os.Open(s)
   if e != nil {
      return "", e
   }
   build_o := strings.Builder{}
   io.Copy(&build_o, open_o)
   return build_o.String(), nil
}

func main() {
   s, e := GetContents("a.go")
   if e != nil {
      os.Exit(1)
   }
   print(s)
}
