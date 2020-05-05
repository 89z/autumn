package main
import (
   "fmt"
   "regexp"
)
func main() {
   s1 := "package"
   // example 1
   s2 := regexp.MustCompile("p.").FindString(s1)
   // example 2
   a1 := regexp.MustCompile("a.").FindAllString(s1, -1)
   // example 3
   a2 := regexp.MustCompile("p(..)").FindStringSubmatch(s1)
   // print
   fmt.Println(s2, a1, a2)
}
