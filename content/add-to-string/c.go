package main
import (
   "bytes"
   "fmt"
)
func main() {
   var o1 bytes.Buffer
   // example 1
   o1.WriteString("Sun")
   // example 2
   fmt.Fprint(&o1, "day")
   // print
   var s1 = o1.String()
   println(s1)
}
