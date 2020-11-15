package main

import (
   "fmt"
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
   y, e := GetContents("a.go")
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s", y)
}
