package main

import (
   "log"
   "os/exec"
   "strings"
   "io"
)

func shellExec(name string, arg ...string) (string, error) {
   cmd := exec.Command(name, arg...)
   out, err := cmd.StdoutPipe()
   if err != nil { return "", err }
   cmd.Start()
   b := new(strings.Builder)
   io.Copy(b, out)
   return strings.TrimRight(b.String(), "\n"), nil
}

func main() {
   s, err := shellExec("go", "version")
   if err != nil {
      log.Fatal(err)
   }
   println(s)
}
