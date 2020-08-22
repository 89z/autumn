package main
import (
   "fmt"
   "os"
   "path/filepath"
)
func main() {
   a, e := filepath.Glob("*")
   if e != nil {
      os.Exit(1)
   }
   fmt.Println(a)
}
