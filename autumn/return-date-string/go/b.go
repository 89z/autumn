package main

import (
   "log"
   "os"
   "time"
)

func main() {
   b, e := time.Now().MarshalText()
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
