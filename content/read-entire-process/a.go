package main
import (
   "fmt"
   "os/exec"
)
func main() {
   t1 := exec.Command("cal")
   a1, _ := t1.Output()
   fmt.Printf("%s", a1)
}
