package main

import (
   "bufio"
   "bytes"
   "log"
   "os/exec"
)

func output(command ...string) (*bufio.Scanner, error) {
   name, arg := command[0], command[1:]
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
