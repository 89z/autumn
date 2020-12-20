package main

import (
   "bytes"
   "log"
   "os"
)

func GetContents(s string) (bytes.Buffer, error) {
   buf_o := bytes.Buffer{}
   open_o, e := os.Open(s)
   if e != nil {
      return buf_o, e
   }
   buf_o.ReadFrom(open_o)
   return buf_o, nil
}

func main() {
   o, e := GetContents("readfrom.go")
   if e != nil {
      log.Fatal(e)
   }
   o.WriteTo(os.Stdout)
}
