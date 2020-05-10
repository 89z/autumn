package main
import (
   "os"
   "os/exec"
)
func main() {
   o1 := exec.Command("less", "-V")
   o1.Stdout = os.Stdout
   o1.Run()
}
