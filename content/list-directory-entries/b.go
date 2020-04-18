package main
import (
   "fmt"
   "os"
)
func main() {
   f1, _ := os.Open(".")
   a1, _ := f1.Readdirnames(0)
   fmt.Println(a1)
}
