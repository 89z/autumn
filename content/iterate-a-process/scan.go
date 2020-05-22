package main
import (
   "bufio"
   "bytes"
   "log"
   "os/exec"
)
func main() {
   y1, e1 := exec.Command("ag", "-V").Output()
   if e1 != nil {
      log.Fatal(e1)
   }
   o1 := bytes.NewReader(y1)
   o2 := bufio.NewScanner(o1)
   for o2.Scan() {
      s1 := o2.Text()
      println(s1)
   }
}
