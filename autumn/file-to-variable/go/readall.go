package main

import (
   "io/ioutil"
   "log"
   "os"
)

func main() {
   o, e := os.Open("readall.go")
   if e != nil {
      log.Fatal(e)
   }
   y, e := ioutil.ReadAll(o)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(y)
}
