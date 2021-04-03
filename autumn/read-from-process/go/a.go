package main

import (
   "log"
   "os/exec"
   "strings"
)

func shellExec(name string, arg ...string) (string, error) {
   out, err := exec.Command(name, arg...).Output()
   if err != nil { return "", err }
   return strings.TrimRight(string(out), "\n"), nil
}

func main() {
   out, err := shellExec("go", "version")
   if err != nil {
      log.Fatal(err)
   }
   println(out)
}
