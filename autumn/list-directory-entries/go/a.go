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
   for _, o := range a {
      s := o.Name()
      println(s)
   }
}
