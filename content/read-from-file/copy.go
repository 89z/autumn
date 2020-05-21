package main
import (
   "io"
   "os"
   "strings"
)
func main() {
   var o1, e1 = os.Open("a.txt")
   if e1 == nil {
      var o2 strings.Builder
      io.Copy(&o2, o1)
      var s1 = o2.String()
      print(s1)
   }
}
