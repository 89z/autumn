package main
import "os"
import "strings"
func main() {
   o1 := strings.NewReader("Sunday\n")
   o1.WriteTo(os.Stdout)
}
