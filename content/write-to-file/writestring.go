package main
import "os"
func main() {
   o1, _ := os.Create("a.txt")
   o1.WriteString("Sunday\n")
}
