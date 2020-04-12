package main
import (
   "os"
   "os/exec"
)
func main() {
   // example 1
   r1 := exec.Command("ag")
   r1.Stdout = os.Stdout
   r1.Run()
   // example 2
   r2 := exec.Command("ag", "-V")
   r2.Stdout = os.Stdout
   r2.Run()
}
