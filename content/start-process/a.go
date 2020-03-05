package main
import (
   "os"
   "os/exec"
)
func main() {
   r1 := exec.Command("cal")
   r1.Stdout = os.Stdout
   r1.Run()
}
