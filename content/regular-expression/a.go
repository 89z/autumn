package main
import (
   "fmt"
   "regexp"
)
func main() {
   s1 := []byte("2020-03-29")
   // example 1
   o1 := regexp.MustCompile("-..")
   s2 := o1.Find(s1)
   // example 2
   o2 := regexp.MustCompile("-..")
   a1 := o2.FindAll(s1, -1)
   // example 3
   o3 := regexp.MustCompile("-(..)-(..)")
   a2 := o3.FindSubmatch(s1)
   // print
   fmt.Printf("%s\n", []interface{}{s2, a1, a2})
}
