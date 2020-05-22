package main
import (
   "bytes"
   "fmt"
   "log"
   "os/exec"
)
func f(r rune) bool {
   return r == '\n'
}
func main() {
   y, e := exec.Command("ag", "-V").Output()
   if e != nil {
      log.Fatal(e)
   }
   a := bytes.FieldsFunc(y, f)
   for n, y := range a {
      fmt.Printf("%v %s\n", n, y)
   }
}
