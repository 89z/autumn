package main

import (
   "os"
   "os/exec"
)

func main() {
   o := exec.Command("git", "status")
   o.Dir = "/Git/winter"
   o.Stdout = os.Stdout
   o.Run()
}
