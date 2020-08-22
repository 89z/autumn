package main
import (
   "fmt"
   "os"
   "os/exec"
)
func main() {
   y, e := exec.Command("go", "version").Output()
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s", y)
}
