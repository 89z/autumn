package main
import (
   "fmt"
   "os/exec"
)
func main() {
   o1 := exec.Command("ag")
   a1, _ := o1.Output()
   fmt.Printf("%s", a1)
}
