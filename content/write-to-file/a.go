package main
import "os"
func main() {
   r1, _ := os.Create("day.txt")
   r1.WriteString("Sunday\n")
}
