package main
import "os/exec"

func main() {
   exec.Command("gvim").Run()
}
