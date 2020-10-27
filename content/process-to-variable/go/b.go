package main

import (
   "os/exec"
   "strings"
)

func main() {
   o := exec.Command("go", "version")
   o1 := strings.Builder{}
   o.Stdout = &o1
   o.Run()
   s := o1.String()
   print(s)
}
