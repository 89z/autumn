package main

import (
   "bufio"
   "bytes"
   "log"
   "os/exec"
)

func shellExec(name string, arg ...string) (string, error) {
   out, err := exec.Command(name, arg...).Output()
   if err != nil { return nil, err }
   return bufio.NewScanner(bytes.NewReader(out)), nil
}

func main() {
   env, err := shellExec("go", "env")
   if err != nil {
      log.Fatal(err)
   }
   for env.Scan() {
      println(env.Text())
   }
}
