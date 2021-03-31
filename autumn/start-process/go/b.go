package main

import (
   "os"
   "os/exec"
)

func main() {
   c := exec.Command("go", "env")
   c.Stdout = os.Stdout
   c.Run()
}
