package main

import (
   "bufio"
   "bytes"
   "log"
   "os/exec"
)

func output(name string, arg ...string) (*bufio.Scanner, error) {
   b, e := exec.Command(name, arg...).Output()
   return bufio.NewScanner(bytes.NewReader(b)), e
}

func main() {
   s, e := output("go", "env")
   if e != nil {
      log.Fatal(e)
   }
   for s.Scan() {
      println(s.Text())
   }
}
