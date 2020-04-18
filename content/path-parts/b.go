package main
import "path/filepath"
func main() {
   s1 := `C:\sunday\monday.tar.xz`
   // example 1
   b1 := filepath.Dir(s1) == `C:\sunday`
   // example 2
   b2 := filepath.Base(s1) == "monday.tar.xz"
   // example 3
   b3 := filepath.Ext(s1) == ".xz"
   // example 4
   s2, s3 := filepath.Split(s1)
   b4 := s2 == `C:\sunday\`
   b5 := s3 == `monday.tar.xz`
   // print
   println(b1, b2, b3, b4, b5)
}
