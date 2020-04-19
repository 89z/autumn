package main
import "fmt"
import "os"
func main() {
   o1, _ := os.Create("a.txt")
   fmt.Fprintln(o1, "Sunday")
}
