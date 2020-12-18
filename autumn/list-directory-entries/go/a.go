package main

import (
   "io/ioutil"
   "log"
)

func main() {
   a, e := ioutil.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for n := range a {
      s := a[n].Name()
      println(s)
   }
}
