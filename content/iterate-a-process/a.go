package main
import (
   "bufio"
   "os/exec"
)
func main() {
   r1 := exec.Command("ag", "-V")
   r2, _ := r1.StdoutPipe()
   r1.Start()
   r3 := bufio.NewScanner(r2)
   for r3.Scan() {
      s1 := r3.Text()
      println(s1)
   }
}
