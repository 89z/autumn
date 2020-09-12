package main

import (
   "fmt"
   "io/ioutil"
   "log"
   "os"
)

func main() {
   o, e := os.Open("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   y, e := ioutil.ReadAll(o)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%s", y)
}
