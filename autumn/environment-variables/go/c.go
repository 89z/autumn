package main

import (
   "log"
   "os"
)

func main() {
   s, e := os.UserCacheDir()
   if e != nil {
      log.Fatal(e)
   }
   println(s == `C:\Users\Steven\AppData\Local`)
}
