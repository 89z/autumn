package main

import (
   "os"
   "os/exec"
)

func system(command ...string) error {
   name, arg := command[0], command[1:]
   c := exec.Command(name, arg...)
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   return c.Run()
}

func main() {
   system("go", "version")
}
