package main
import "os"
func main() {
   os.Symlink("a.txt", "b.txt")
}
