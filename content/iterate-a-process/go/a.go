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
   o1 := bufio.NewScanner(o)
   for o1.Scan() {
      s := o1.Text()
      println(s)
   }
}
