package main

import (
   "bufio"
   "log"
   "os/exec"
)

func popen(name string, arg ...string) (*bufio.Scanner, error) {
   cmd := exec.Command(name, arg...)
   pipe, e := cmd.StdoutPipe()
   if e != nil {
      return nil, e
   }
   return bufio.NewScanner(pipe), cmd.Start()
}

func main() {
   env, e := popen("go", "env")
   if e != nil {
      log.Fatal(e)
   }
   for env.Scan() {
      println(env.Text())
   }
}
