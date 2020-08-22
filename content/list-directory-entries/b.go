package main
import (
   "fmt"
   "os"
)
func main() {
   o, e := os.Open(".")
   if e != nil {
      os.Exit(1)
   }
   a, e := o.Readdirnames(0)
   if e != nil {
      os.Exit(1)
   }
   fmt.Println(a)
}
