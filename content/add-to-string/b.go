package main
import (
   "fmt"
   "strings"
)
func main() {
   var o1 strings.Builder
   // example 1
   o1.WriteString("Sun")
   // example 2
   fmt.Fprint(&o1, "day")
   // print
   var s1 = o1.String()
   println(s1)
}
