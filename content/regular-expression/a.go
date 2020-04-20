package main
import (
   "fmt"
   "regexp"
)
func main() {
   s1 := "2019-12-31"
   // example 1
   o1 := regexp.MustCompile(`\d`)
   b1 := o1.MatchString(s1)
   // example 2
   o2 := regexp.MustCompile("-..")
   s2 := o2.FindString(s1)
   // example 3
   o3 := regexp.MustCompile("-..")
   a1 := o3.FindAllString(s1, -1)
   // example 4
   o4 := regexp.MustCompile("-(..)-(..)")
   a2 := o4.FindStringSubmatch(s1)
   // print
   fmt.Println(b1, s2, a1, a2)
}
