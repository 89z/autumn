package main
import (
   "fmt"
   "strings"
)
func main() {
   s1 := "sun mon tue"
   // example 1
   a1 := strings.Split(s1, " ")
   // example 2
   a2 := strings.SplitN(s1, " ", 2)
   // example 3
   a3 := strings.Fields(s1)
   // print
   fmt.Printf("%q\n", a1)
   fmt.Printf("%q\n", a2)
   fmt.Printf("%q\n", a3)
}
