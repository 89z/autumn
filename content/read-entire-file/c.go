package main
import (
   "io"
   "os"
   "strings"
)
func main() {
   // read
   var o1, _ = os.Open("a.txt")
   var o2 strings.Builder
   io.Copy(&o2, o1)
   // write
   var s1 = o2.String()
   print(s1)
}
