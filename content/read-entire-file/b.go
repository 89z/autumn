package main
import (
   "bytes"
   "io"
   "os"
)
func main() {
   var o1, _ = os.Open("a.txt")
   var o2 bytes.Buffer
   io.Copy(&o2, o1)
   // print
   var s1 = o2.String()
   print(s1)
}
