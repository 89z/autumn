package main

import (
   "bufio"
   "bytes"
   "log"
   "os/exec"
)

func main() {
   y, e := exec.Command("go", "env").Output()
   if e != nil {
      log.Fatal(e)
   }
   o := bufio.NewScanner(bytes.NewReader(y))
   for o.Scan() {
      s := o.Text()
      println(s)
   }
}
