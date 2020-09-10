package main
import (
   "fmt"
   "io/ioutil"
   "os"
)
func main() {
   y, e := ioutil.ReadFile("a.txt")
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s", y)
}
