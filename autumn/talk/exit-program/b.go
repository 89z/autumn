package main

import (
   "fmt"
   "log"
)

func main() {
   e := fmt.Errorf("%v", 9)
   if e != nil {
      log.Fatal(e)
   }
}
