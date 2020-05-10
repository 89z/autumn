package main
import (
   "bytes"
   "fmt"
   "os/exec"
)
func f1(r1 rune) bool {
   return r1 == '\n'
}
func main() {   
   y1, _ := exec.Command("ag", "-V").Output()
   a1 := bytes.FieldsFunc(y1, f1)
   for _, s1 := range a1 {
      fmt.Printf("%s\n", s1)
   }
}
