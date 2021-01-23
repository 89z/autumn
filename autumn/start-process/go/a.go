package main

import (
   "os"
   "os/exec"
)

func system(name string, arg ...string) error {
   cmd := exec.Command(name, arg...)
   cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
   return cmd.Run()
}

func main() {
   system("go", "version")
}
