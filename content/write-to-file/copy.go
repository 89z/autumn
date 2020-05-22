package main
import (
   "io"
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
   io.Copy(o2, o1)
}
