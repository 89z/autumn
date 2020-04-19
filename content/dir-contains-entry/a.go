package main
import (
   "fmt"
   "os"
)
func main() {
   s1 := "index.md"
   // example 1
   _, e1 := os.Stat(s1)
   // example 2
   _, e2 := os.Stat(s1)
   b1 := os.IsNotExist(e2)
   // example 3
   o1, _ := os.Stat(s1)
   b2 := o1.Mode().IsRegular()
   // print
   fmt.Println(e1, b1, b2)
}
