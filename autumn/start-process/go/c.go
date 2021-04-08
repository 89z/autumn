package main

import (
   "os"
   "os/exec"
)

func main() {
   c := exec.Command("dust")
   c.Stdout = os.Stdout
   c.Run()
}
