package main
import "os/exec"

func main() {
   c := exec.Command("go", "mod", "init", "day")
   c.Dir = `C:\Users\Public`
   c.Run()
}
