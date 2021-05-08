package main

import (
   "os"
   "os/exec"
)

func main() {
   c := exec.Command("dust", "-V")
   p, e := c.StdoutPipe()
   if e != nil {
      panic(e)
   }
   c.Start()
   os.Stdout.ReadFrom(p)
   defer c.Wait()
}
