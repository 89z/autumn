package main
import (
   "os"
   "os/exec"
)
func main() {
   o2 := exec.Command("less", "index.md")
   o2.Stdin = os.Stdin
   o2.Stdout = os.Stdout
   o2.Start()
}
