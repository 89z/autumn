package main

import (
   "io/ioutil"
   "os"
)

func GetContents(s string) ([]byte, error) {
   o, e := os.Open(s)
   if e != nil {
      return []byte{}, e
   }
   return ioutil.ReadAll(o)
}

func main() {
   GetContents("100mb.file")
}
