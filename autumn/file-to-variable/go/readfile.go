package main

import (
   "io/ioutil"
   "log"
   "os"
)

func main() {
   y, e := ioutil.ReadFile("readfile.go")
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(y)
}
