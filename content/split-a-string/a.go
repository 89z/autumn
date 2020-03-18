package main
import (
   "fmt"
   "strings"
)
func main() {
   s1 := "Sun Mon Tue"
   // example 1
   a1 := strings.Split(s1, " ")
   // example 2
   a2 := strings.SplitN(s1, " ", 2)
   // example 3
   a3 := strings.Fields(s1)
   // print
   fmt.Printf("%q\n", [][]string{a1, a2, a3})
}
