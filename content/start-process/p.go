package main

import (
   "os"
   "os/exec"
)

func main() {
   println("run")
   o := exec.Command("pipe")
   o.Stdout = os.Stdout
   o.Run()
}
