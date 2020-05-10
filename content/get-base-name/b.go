package main
import "path/filepath"
func main() {
   s1 := `C:\sunday\monday.tar.xz`
   // example 1
   s2 := filepath.Base(s1)
   // example 2
   _, s3 := filepath.Split(s1)
   // example 3
   s4 := filepath.Ext(s1)
   // print
   println(s2 == "monday.tar.xz", s3 == "monday.tar.xz", s4 == ".xz")
}
