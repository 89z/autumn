package main
import (
   "os"
   "path/filepath"
)
func main() {
   s1 := "index.md"
   s2, e := filepath.Abs(s1)
   if e != nil {
      os.Exit(1)
   }
   println(s2)
}
