package main
import (
   "bufio"
   "os/exec"
)
func main() {
   o1 := exec.Command("ag", "-V")
   o2, _ := o1.StdoutPipe()
   o1.Start()
   s1, _ := bufio.NewReader(o2).ReadString('\n')
   print(s1)
}
