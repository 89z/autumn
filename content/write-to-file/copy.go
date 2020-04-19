package main
import "io"
import "os"
import "strings"
func main() {
   o1 := strings.NewReader("Sunday\n")
   o2, _ := os.Create("a.txt")
   io.Copy(o2, o1)
}
