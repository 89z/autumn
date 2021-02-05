package main

import (
   "os"
   "os/exec"
)

func main() {
   c := exec.Command("git", "status")
   c.Dir = "/Git/autumn"
   c.Stdout = os.Stdout
   c.Run()
}
