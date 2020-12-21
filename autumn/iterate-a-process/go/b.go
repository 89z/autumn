package main

import (
   "bytes"
   "log"
   "os"
   "os/exec"
)

func f(r rune) bool {
   return r == '\n'
}

func main() {
   y, e := exec.Command("go", "env").Output()
   if e != nil {
      log.Fatal(e)
   }
   a := bytes.FieldsFunc(y, f)
   for _, y := range a {
      y = append(y, '\n')
      os.Stdout.Write(y)
   }
}
