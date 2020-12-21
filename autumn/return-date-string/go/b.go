package main

import (
   "log"
   "os"
   "time"
)

func main() {
   y, e := time.Now().MarshalText()
   if e != nil {
      log.Fatal(e)
   }
   y = append(y, '\n')
   os.Stdout.Write(y)
}
