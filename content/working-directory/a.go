package main
import (
   "log"
   "os"
)
func main() {
   s, e := os.Getwd()
   if e != nil {
      log.Fatal(e)
   }
   println(s)
}
