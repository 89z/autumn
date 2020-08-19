package main
import (
   "fmt"
   "log"
   "os"
)
func main() {
   o, e := os.Create("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Fprintln(o, "Sunday")
}
