package main

import (
   "os"
   "os/exec"
)

func main() {
   c := exec.Command("dust")
   c.Dir, c.Stdout = "..", os.Stdout
   c.Run()
}
