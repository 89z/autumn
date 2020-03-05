package main
import (
   "bufio"
   "os/exec"
)
func main() {
   r1 := exec.Command("cal")
   r2, _ := r1.StdoutPipe()
   r1.Start()
   r3 := bufio.NewScanner(r2)
   for r3.Scan() {
      println(r3.Text())
   }
}
