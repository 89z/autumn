package main
import "os/exec"

func main() {
   o := exec.Command("waterfox", "google.com/search?tbm=vid&q=squarepusher")
   o.Run()
}
