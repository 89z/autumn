package main
import (
   "fmt"
   "path"
)
func main() {
   s1 := "/sunday/monday.txt"
   // example 1
   s2 := path.Dir(s1)
   // example 2
   s3 := path.Base(s1)
   // print
   fmt.Printf("%q\n", []string{s2, s3})
}
