package main
import (
   "fmt"
   "strings"
)
func main() {
   // example 1
   var s1 = "Sun"
   s1 = s1 + "day"
   // example 2
   var s2 = "Sun"
   s2 += "day"
   // example 3
   var s3 = "Sun"
   s3 = fmt.Sprintf("%s Mon", s3)
   // example 4
   var s4 strings.Builder
   s4.WriteString("Sun")
   s4.WriteString("day")
   // print
   fmt.Printf("%q\n", []string{s1, s2, s3, s4.String()})
}
