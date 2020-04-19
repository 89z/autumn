package main
import (
   "bytes"
   "os"
)
func main() {
   // in
   var o1 bytes.Buffer
   o1.WriteString("Sunday\n")
   // out
   var o2, _ = os.Create("a.txt")
   o1.WriteTo(o2)
}
