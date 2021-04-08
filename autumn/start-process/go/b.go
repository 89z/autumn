package main

import (
   "bytes"
   "os/exec"
)

func main() {
   b, e := exec.Command(name, arg...).Output()
   if e != nil { return nil, e }
   return bytes.TrimRight(b, "\n"), nil
   b, e := shellExec("go", "version")
   if e != nil {
      panic(e)
   }
   println(string(b))
}
