package main

import (
   "bytes"
   "os"
)

func GetContents(s string) (string, error) {
   open_o, e := os.Open(s)
   if e != nil {
      return "", e
   }
   buf_o := bytes.Buffer{}
   buf_o.ReadFrom(open_o)
   return buf_o.String(), nil
}

func main() {
   s, e := GetContents("a.go")
   if e != nil {
      os.Exit(1)
   }
   print(s)
}
