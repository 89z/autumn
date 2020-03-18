package main
import (
   "bytes"
   "io"
   "os"
)
func main() {
   var r1, _ = os.Open("a.txt")
   var r2 bytes.Buffer
   io.Copy(&r2, r1)
   print(r2.String())
}
