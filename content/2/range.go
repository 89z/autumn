package main
import (
   "bytes"
   "fmt"
   "os"
   "os/exec"
)
func f(r rune) bool {
   return r == '\n'
}
func main() {
   y, e := exec.Command("go", "env").Output()
   if e != nil {
      os.Exit(1)
   }
   a := bytes.FieldsFunc(y, f)
   for n, y := range a {
      fmt.Printf("%v %s\n", n, y)
   }
}
