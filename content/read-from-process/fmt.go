package main
import (
   "fmt"
   "log"
   "os/exec"
)
func main() {
   o := exec.Command("ag", "-V")
   y, e := o.Output()
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%s", y)
}
