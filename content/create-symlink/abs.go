package main
import (
   "os"
   "path/filepath"
)
func main() {
   var s_pub = os.Getenv("PUBLIC")
   var s_old, s_new string
   // round 1
   s_old, _ = filepath.Abs("abs.go")
   s_new = filepath.Join(s_pub, "Music", "abs.go")
   os.Symlink(s_old, s_new)
   // round 2
   s_old, _ = filepath.Abs("chdir.go")
   s_new = filepath.Join(s_pub, "Videos", "chdir.go")
   os.Symlink(s_old, s_new)
}