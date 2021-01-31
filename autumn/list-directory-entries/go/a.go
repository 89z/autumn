package main

import (
   "io/ioutil"
   "log"
)

func main() {
   infos, e := ioutil.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for _, info := range infos {
      s := info.Name()
      println(s)
   }
}
