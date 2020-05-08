package main
import (
   "os"
   "path/filepath"
)
func main() {
   var s_pub = os.Getenv("PUBLIC")
   var s_old, s_new string
   // round 1
   s_old, _ = filepath.Abs("a.go")
   s_new = filepath.Join(s_pub, "Music", "a.go")
   os.Symlink(s_old, s_new)
   // round 2
   s_old, _ = filepath.Abs("b.go")
   s_new = filepath.Join(s_pub, "Videos", "b.go")
   os.Symlink(s_old, s_new)
}
