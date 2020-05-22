package main
import (
   "log"
   "os"
   "path/filepath"
)
func main() {
   s_pub := os.Getenv("PUBLIC")
   s_old, e := filepath.Abs("abs.go")
   if e != nil {
      log.Fatal(e)
   }
   s_new := filepath.Join(s_pub, "Music", "abs.go")
   os.Symlink(s_old, s_new)
}
