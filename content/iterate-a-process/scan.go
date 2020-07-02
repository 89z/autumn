package main
import (
   "bufio"
   "bytes"
   "log"
   "os/exec"
)
func main() {
   y, e := exec.Command("ag", "-V").Output()
   if e != nil {
      log.Fatal(e)
   }
   o1 := bytes.NewReader(y)
   o2 := bufio.NewScanner(o1)
   for o2.Scan() {
      s := o2.Text()
      println(s)
   }
}
