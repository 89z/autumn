package main

import (
   "fmt"
   "io/ioutil"
   "os"
)

func main() {
   y, e := ioutil.ReadFile("a.go")
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s", y)
}
