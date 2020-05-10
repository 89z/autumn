package main
import (
   "fmt"
   "strings"
)
func main() {
   s1 := "Sunday Monday"
   // example 1
   a1 := strings.Split(s1, " ")
   // example 2
   a2 := strings.Fields(s1)
   // print
   fmt.Println(a1, a2)
}
