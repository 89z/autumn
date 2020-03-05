package main
import (
   "io"
   "os"
   "strings"
)
func main() {
   var r1, _ = os.Open("a.txt")
   var s1 strings.Builder
   io.Copy(&s1, r1)
   print(s1.String())
}
