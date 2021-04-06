package main

import (
   "os"
   "time"
)

func main() {
   b, e := time.Now().MarshalText()
   if e != nil {
      panic(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
