package main
import (
   "log"
   "os"
)
func main() {
   s_old, e := os.Getwd()
   if e != nil {
      log.Fatal(e)
   }
   s_new := os.Getenv("TMP")
   os.Symlink(s_old + "/index.md", s_new + "/index.md")
}
