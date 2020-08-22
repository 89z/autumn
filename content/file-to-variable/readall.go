package main
import (
   "fmt"
   "io/ioutil"
   "os"
)
func main() {
   o, e := os.Open("a.txt")
   if e != nil {
      os.Exit(1)
   }
   y, e := ioutil.ReadAll(o)
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s", y)
}
