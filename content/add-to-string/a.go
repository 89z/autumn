package main
import (
   "fmt"
   "strings"
)
func main() {
   // example 1
   var s1 = "sun"
   s1 = s1 + " mon"
   // example 2
   var s2 = "sun"
   s2 += " mon"
   // example 3
   var s3 = "sun"
   s3 = fmt.Sprintf("%s mon", s3)
   // example 4
   var s4 strings.Builder
   s4.WriteString("sun")
   s4.WriteString(" mon")
   // print
   fmt.Printf("%q\n", []string{s1, s2, s3, s4.String()})
}
