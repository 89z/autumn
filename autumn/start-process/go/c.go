package main

import (
   "os/exec"
   "strings"
)

func main() {
   c, b := exec.Command("dust", "-V"), new(strings.Builder)
   c.Stdout = b
   c.Run()
   println(b.String() == "Dust 0.5.4\n")
}
