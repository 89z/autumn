package main
import "os/exec"

func main() {
   b, e := exec.Command("dust", "-V").Output()
   if e != nil {
      panic(e)
   }
   println(string(b) == "Dust 0.5.4\n")
}
