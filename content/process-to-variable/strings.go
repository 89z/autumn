package main
import "os/exec"
import "strings"
func main() {
   var o1 = exec.Command("ag")
   var o2 strings.Builder
   o1.Stdout = &o2
   o1.Run()
   var s1 = o2.String()
   print(s1)
}
