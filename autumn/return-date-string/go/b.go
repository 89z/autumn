package main

import (
   "fmt"
   "log"
   "time"
)

func main() {
   y, e := time.Now().MarshalText()
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%s\n", y)
}
