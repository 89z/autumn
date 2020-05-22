package main
import (
   "log"
   "os"
)
func main() {
   o, e := os.Stat("index.md")
   if os.IsNotExist(e) {
      log.Fatal(e)
   }
   log.Print(o)
}
