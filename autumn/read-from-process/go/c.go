package main

import (
   "bufio"
   "os/exec"
)

func popen(name string, arg ...string) (*bufio.Scanner, error) {
   c := exec.Command(name, arg...)
   p, e := c.StdoutPipe()
   if e != nil { return nil, e }
   return bufio.NewScanner(p), c.Start()
}

func main() {
   s, e := popen("go", "env")
   if e != nil {
      panic(e)
   }
   for s.Scan() {
      println(s.Text())
   }
}
