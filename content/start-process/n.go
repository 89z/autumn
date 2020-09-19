package main
import "os/exec"

func main() {
   o := exec.Command("notepad", "start.go")
   o.Start()
}
