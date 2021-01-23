package main

import (
   "os"
   "os/exec"
)

func system(name string, arg ...string) error {
   c := exec.Command(name, arg...)
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   return c.Run()
}

func main() {
   system("go", "version")
}
