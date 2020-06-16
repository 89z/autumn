package main
import (
   "os"
   "os/exec"
)
func main() {
   o := exec.Command("less", "index.md")
   o.Stdin = os.Stdin
   o.Stdout = os.Stdout
   o.Start()
}
