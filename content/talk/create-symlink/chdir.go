package main
import "os"
func main() {
   s_wd := os.Getwd()
   s_pub := os.Getenv("PUBLIC")
   s_sep := string(os.PathSeparator)
   // round 1
   os.Chdir(s_pub + s_sep + "Music")
   os.Symlink(s_wd + s_sep + "abs.go", "abs.go")
   // round 2
   os.Chdir(s_pub + s_sep + "Videos")
   os.Symlink(s_wd + s_sep + "chdir.go", "chdir.go")
}
