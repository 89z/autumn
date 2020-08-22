package main
import (
   "fmt"
   "os"
)
func main() {
   o, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   fmt.Fprintln(o, "Sunday")
}
