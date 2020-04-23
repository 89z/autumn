package main
import "fmt"
import "os/exec"
func main() {
   o1 := exec.Command("ag")
   y1, _ := o1.Output()
   fmt.Printf("%s", y1)
}
