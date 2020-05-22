package main
import (
   "io/ioutil"
   "log"
)
func main() {
   a, e := ioutil.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for n, o := range a {
      s := o.Name()
      println(n, s)
   }
}
