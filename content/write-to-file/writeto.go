package main
import (
   "log"
   "os"
   "strings"
)
func main() {
   o1 := strings.NewReader("Sunday\n")
   o2, e1 := os.Create("a.txt")
   if e1 != nil {
      log.Fatal(e1)
   }
   o1.WriteTo(o2)
}
