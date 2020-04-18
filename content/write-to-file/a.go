package main
import "os"
func main() {
   f1, _ := os.Create("day.txt")
   f1.WriteString("Sunday\n")
}
