package main
import (
   "log"
   "os"
)
func main() {
   o, e := os.Stat("index.md")
   if e != nil {
      log.Fatal(e)
   }
   log.Print(o)
}
