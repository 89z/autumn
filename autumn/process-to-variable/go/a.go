package main

import (
   "log"
   "os"
   "os/exec"
)

func main() {
   y, e := exec.Command("go", "version").Output()
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(y)
}
