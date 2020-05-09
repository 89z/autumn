package main
import (
   "os"
   "path/filepath"
)
func main() {
   var s_pub = os.Getenv("PUBLIC")
   var s_sep = string(os.PathSeparator)
   var s_old, s_new string
   // round 1
   s_old, _ = filepath.Abs("abs.go")
   os.Symlink(s_old, s_pub + s_sep + "Music" + s_sep + "abs.go")
   // round 2
   s_old, _ = filepath.Abs("chdir.go")
   os.Symlink(s_old, s_pub + s_sep + "Music" + s_sep + "chdir.go")
}
