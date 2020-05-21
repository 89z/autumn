package main
import "io"
import "os/exec"
import "strings"
func main() {
   var o1 = exec.Command("ag", "-V")
   var o2, _ = o1.StdoutPipe()
   o1.Start()
   var o3 strings.Builder
   io.Copy(&o3, o2)
   var s1 = o3.String()
   print(s1)
}
