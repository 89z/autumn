package main
import (
   "bufio"
   "os/exec"
)
func main() {
   p1 := exec.Command("cal")
   p2, _ := p1.StdoutPipe()
   p1.Start()
   s1, _ := bufio.NewReader(p2).ReadString('\n')
   print(s1)
}
