package main
import (
   "fmt"
   "regexp"
)
func main() {
   s1 := "Wednesday"
   s2 := "e(..)"
   // example 1
   a1 := regexp.MustCompile(s2).FindStringSubmatch(s1)
   // example 2
   a2 := regexp.MustCompile(s2).FindAllStringSubmatch(s1, -1)
   // print
   fmt.Println(a1, a2)
}
