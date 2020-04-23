package main
import (
   "os"
   "os/exec"
)
func main() {
   // example 1
   o1 := exec.Command("ag")
   o1.Stdout = os.Stdout
   o1.Run()
   // example 2
   o2 := exec.Command("ag", "-V")
   o2.Stdout = os.Stdout
   o2.Run()
}
