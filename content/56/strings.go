package main

import (
   "os/exec"
   "strings"
)

func main() {
   o := exec.Command("go", "version")
   o2 := strings.Builder{}
   o.Stdout = &o2
   o.Run()
   s := o2.String()
   print(s)
}
