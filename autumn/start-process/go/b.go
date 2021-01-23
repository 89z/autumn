package main

import (
   "os"
   "os/exec"
)

func main() {
   cmd := exec.Command("go", "env")
   cmd.Stdout = os.Stdout
   cmd.Start()
   cmd.Wait()
}
