package main
import (
   "io/ioutil"
   "os"
)
func main() {
   a, e := ioutil.ReadDir(".")
   if e != nil {
      os.Exit(1)
   }
   for n, o := range a {
      s := o.Name()
      println(n, s)
   }
}
