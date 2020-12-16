package main

import (
   "io/ioutil"
   "time"
)

func main() {
   ioutil.ReadFile("100mb.file")
   time.Sleep(time.Duration(time.Minute))
}
