package main
import (
   "fmt"
   "regexp"
)
func main() {
   s1 := "Wednesday"
   // example 1
   s2 := regexp.MustCompile("e.").FindString(s1)
   // example 2
   a1 := regexp.MustCompile("e.").FindAllString(s1, -1)
   // example 3
   a2 := regexp.MustCompile("e(..)").FindStringSubmatch(s1)
   // example 4
   a3 := regexp.MustCompile("e(..)").FindAllStringSubmatch(s1, -1)
   // print
   fmt.Println(s2, a1, a2, a3)
}
