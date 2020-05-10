package main
import (
   "bufio"
   "bytes"
   "os/exec"
)
func main() {   
   y1, _ := exec.Command("ag", "-V").Output()
   o1 := bytes.NewReader(y1)
   o2 := bufio.NewScanner(o1)
   for o2.Scan() {
      s1 := o2.Text()
      println(s1)
   }
}
