package main

import (
   "log"
   "os/exec"
   "strings"
)

func shellExec(name string, arg ...string) (string, error) {
   c, b := exec.Command(name, arg...), new(strings.Builder)
   c.Stdout = b
   e := c.Run()
   return strings.TrimRight(b.String(), "\n"), e
}

func main() {
   s, e := shellExec("go", "version")
   if e != nil {
      log.Fatal(e)
   }
   println(s)
}
