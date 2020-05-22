package main
import (
   "log"
   "os"
)
func main() {
   f, e := os.Open(".")
   if e != nil {
      log.Fatal(e)
   }
   a, e := f.Readdirnames(0)
   if e != nil {
      log.Fatal(e)
   }
   log.Println(a)
}
