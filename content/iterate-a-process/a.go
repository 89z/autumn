package main
import (
   "bufio"
   "os/exec"
)
func main() {
   o1 := exec.Command("ag", "-V")
   o2, _ := o1.StdoutPipe()
   o1.Start()
   o3 := bufio.NewScanner(o2)
   for o3.Scan() {
      s1 := o3.Text()
      println(s1)
   }
}
