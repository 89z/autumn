package main

import (
   "bufio"
   "bytes"
   "os"
   "os/exec"
)

func main() {
   y, e := exec.Command("go", "env").Output()
   if e != nil {
      os.Exit(1)
   }
   o := bytes.NewReader(y)
   o2 := bufio.NewScanner(o)
   for o2.Scan() {
      s := o2.Text()
      println(s)
   }
}
