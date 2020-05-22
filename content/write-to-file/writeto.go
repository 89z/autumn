package main
import (
   "log"
   "os"
   "strings"
)
func main() {
   o1 := strings.NewReader("Sunday\n")
   o2, e := os.Create("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   o1.WriteTo(o2)
}
