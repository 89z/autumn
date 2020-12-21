package main

import (
   "bytes"
   "os"
   "os/exec"
)

func main() {
   in_o := exec.Command("go", "version")
   out_o := bytes.Buffer{}
   in_o.Stdout = &out_o
   in_o.Run()
   out_o.WriteTo(os.Stdout)
}
