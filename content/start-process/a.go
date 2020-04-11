package main
import (
   "os"
   "os/exec"
)
func main() {
   r1 := exec.Command("ag", "-V")
   r1.Stdout = os.Stdout
   r1.Run()
}
