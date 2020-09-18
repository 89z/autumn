package main

import (
   "os/exec"
   "strings"
)

func main() {
   oA := exec.Command("go", "version")
   oB := strings.Builder{}
   oA.Stdout = &oB
   oA.Run()
   s := oB.String()
   print(s)
}
