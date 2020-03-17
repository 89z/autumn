package main
import (
   "fmt"
   "strings"
)
func main() {
   s1 := " Sun Mon "
   // example 1
   s2 := strings.TrimLeft(s1, " ")
   // example 2
   s3 := strings.TrimRight(s1, " ")
   // example 3
   s4 := strings.Trim(s1, " ")
   // example 4
   s5 := strings.TrimSpace(s1)
   // print
   fmt.Printf("%q\n", []string{s2, s3, s4, s5})
}
