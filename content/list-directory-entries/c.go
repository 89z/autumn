package main
import (
   "fmt"
   "path/filepath"
)
func main() {
   a1, _ := filepath.Glob("*")
   fmt.Println(a1)
}
