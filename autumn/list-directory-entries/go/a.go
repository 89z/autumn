package main

import (
   "io/ioutil"
   "log"
)

func main() {
   dir, e := ioutil.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for _, each := range dir {
      println(each.Name())
   }
}
