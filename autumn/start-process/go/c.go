package main

import (
   "os/exec"
   "strings"
)

func shellExec(name string, arg ...string) (string, error) {
   c, b := exec.Command(name, arg...), new(strings.Builder)
   c.Stdout = b
   e := c.Run()
   return strings.TrimRight(b.String(), "\n"), e
}

func main() {
   s, e := shellExec("go", "version")
   if e != nil {
      panic(e)
   }
   println(s)
}
package main

import (
   "os"
   "os/exec"
)

func main() {
   c := exec.Command("dust")
   c.Stdout = os.Stdout
   c.Run()
}
