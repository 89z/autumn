package main
import (
   "fmt"
   "regexp"
)
func main() {
   s1 := []byte("2020-03-29")
   // example 1
   r1 := regexp.MustCompile("-..")
   s2 := r1.Find(s1)
   // example 2
   r2 := regexp.MustCompile("-..")
   a1 := r2.FindAll(s1, -1)
   // example 3
   r3 := regexp.MustCompile("-(..)-(..)")
   a2 := r3.FindSubmatch(s1)
   // print
   fmt.Printf("%s\n", []interface{}{s2, a1, a2})
}
