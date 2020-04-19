package main
import (
   "fmt"
   "os"
)
func main() {
   s1 := "a.txt"
   // example 1
   _, e1 := os.Stat(s1)
   fmt.Println(e1)
   // example 2
   _, e2 := os.Stat(s1)
   b1 := os.IsNotExist(e2)
   println(b1)
}
