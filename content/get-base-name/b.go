package main
import "path/filepath"
func main() {
   s1 := `C:\sunday\monday.tar.xz`
   // example 1
   b1 := filepath.Base(s1) == "monday.tar.xz"
   // example 2
   b2 := filepath.Ext(s1) == ".xz"
   // example 3
   _, s2 := filepath.Split(s1)
   b3 := s2 == "monday.tar.xz"
   // print
   println(b1, b2, b3)
}
