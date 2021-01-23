package main

import (
   "bufio"
   "log"
   "os/exec"
)

func popen(command ...string) (*bufio.Scanner, error) {
   name, arg := command[0], command[1:]
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
