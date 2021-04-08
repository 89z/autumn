package main

import (
   "bytes"
   "os/exec"
)

func shellExec(name string, arg ...string) ([]byte, error) {
   b, e := exec.Command(name, arg...).Output()
   if e != nil { return nil, e }
   return bytes.TrimRight(b, "\n"), nil
}

func main() {
   b, e := shellExec("go", "version")
   if e != nil {
      panic(e)
   }
   println(string(b))
}
