package main
import "path/filepath"
func main() {
   s1 := `C:\sunday\monday.tar.xz`
   // example 1
   s2 := filepath.Dir(s1)
   // example 2
   s3, _ := filepath.Split(s1)
   // print
   println(s2 == `C:\sunday`, s3 == `C:\sunday\`)
}
