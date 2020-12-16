package main

import (
   "io/ioutil"
   "os"
   "time"
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
   time.Sleep(time.Duration(time.Minute))
}
