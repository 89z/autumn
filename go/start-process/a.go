package main

import (
   "os"
   "os/exec"
)

func main() {
   println("BEGIN")
   o := exec.Command("pipe")
   o.Stdout = os.Stdout
   o.Run()
}
