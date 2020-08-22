package main
import (
   "os/exec"
   "strings"
)
func main() {
   o1 := exec.Command("go", "version")
   o2 := strings.Builder{}
   o1.Stdout = &o2
   o1.Run()
   s := o2.String()
   print(s)
}
