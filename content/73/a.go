package main

import (
   "os"
   "os/exec"
)

func main() {
   o := exec.Command("less", "-V")
   o.Stdout = os.Stdout
   o.Run()
}
