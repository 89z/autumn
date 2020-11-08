package main
import "os/exec"

func main() {
   o := exec.Command("waterfox", "google.com/search?q=squarepusher")
   o.Start()
}
