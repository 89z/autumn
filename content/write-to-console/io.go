package main
import "io"
import "os"
import "strings"
func main() {
   o1 := strings.NewReader("Sunday\n")
   io.Copy(os.Stdout, o1)
}
