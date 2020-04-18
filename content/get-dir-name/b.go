package main
import "path/filepath"
func main() {
   s1 := `C:\sunday\monday.tar.xz`
   // example 1
   b1 := filepath.Dir(s1) == `C:\sunday`
   // example 2
   s2, _ := filepath.Split(s1)
   b2 := s2 == `C:\sunday\`
   // print
   println(b1, b2)
}
