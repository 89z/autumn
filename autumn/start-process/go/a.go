package main

import (
   "os"
   "os/exec"
)

func System(command ...string) error {
   name, arg := command[0], command[1:]
   o := exec.Command(name, arg...)
   o.Stdout = os.Stdout
   return o.Run()
}

func main() {
   System("go", "version")
}
