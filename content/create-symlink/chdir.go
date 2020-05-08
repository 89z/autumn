package main
import (
   "os"
   "path/filepath"
)
func main() {
   var s_wd = os.Getwd()
   var s_pub = os.Getenv("PUBLIC")
   var s_old, s_new string
   // round 1
   s_old = filepath.Join(s_wd, "a.go")
   s_new = filepath.Join(s_pub, "Music")
   os.Chdir(s_new)
   os.Symlink(s_old, "a.go")
   // round 2
   s_old = filepath.Join(s_wd, "b.go")
   s_new = filepath.Join(s_pub, "Videos")
   os.Chdir(s_new)
   os.Symlink(s_old, "b.go")
}
