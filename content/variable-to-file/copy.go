package main
import (
   "io"
   "os"
   "strings"
)
func main() {
   o1 := strings.NewReader("Sunday\n")
   o2, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   io.Copy(o2, o1)
}
