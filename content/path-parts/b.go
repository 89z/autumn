package main
import "path/filepath"
func main() {
   s := `C:\a\b.tar.xz`
   // example 1
   s1, s2 := filepath.Split(s)
   // example 2
   s3 := filepath.Base(s)
   // example 3
   s4 := filepath.Ext(s)
   // print
   println(s1 == `C:\a\`, s2 == "b.tar.xz", s3 == "b.tar.xz", s4 == ".xz")
}
