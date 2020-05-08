package main
import (
   "os"
   "path/filepath"
)
func main() {
   s1, _ := filepath.Abs("index.md")
   s2 := os.Getenv("HOMEDRIVE")
   // example 1
   s3 := filepath.Join(s2, "a.md")
   os.Symlink(s1, s3)
   // example 2
   os.Chdir(s2)
   os.Symlink(s1, "b.md")
}
