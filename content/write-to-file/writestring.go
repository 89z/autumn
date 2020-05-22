package main
import (
   "log"
   "os"
)
func main() {
   o, e := os.Create("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   o.WriteString("Sunday\n")
}
