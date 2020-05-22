package main
import (
   "io"
   "log"
   "os"
   "strings"
)
func main() {
   var o1, e = os.Open("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   var o2 strings.Builder
   io.Copy(&o2, o1)
   var s1 = o2.String()
   print(s1)
}
