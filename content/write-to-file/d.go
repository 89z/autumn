package main
import (
   "os"
   "strings"
)
func main() {
   o1 := strings.NewReader("Sunday\n")
   o2, _ := os.Create("a.txt")
   o1.WriteTo(o2)
}
