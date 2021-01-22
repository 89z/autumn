package main
import "os/exec"

func system(command ...string) error {
   name, arg := command[0], command[1:]
   return exec.Command(name, arg...).Start()
}

func main() {
   system("firefox", "google.com/search?q=squarepusher")
}
