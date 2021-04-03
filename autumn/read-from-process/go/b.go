package main

import (
   "bufio"
   "log"
   "os/exec"
)

func popen(name string, arg ...string) (*bufio.Scanner, error) {
   cmd := exec.Command(name, arg...)
   out, err := cmd.StdoutPipe()
   if err != nil { return nil, err }
   return bufio.NewScanner(out), cmd.Start()
}

func main() {
   env, err := popen("go", "env")
   if err != nil {
      log.Fatal(err)
   }
   for env.Scan() {
      println(env.Text())
   }
}
