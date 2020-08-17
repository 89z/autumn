package main
import (
   "io/ioutil"
   "log"
)
func main() {
   y, e := ioutil.ReadFile("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   log.Printf("%s", y)
}
