package main
import "path/filepath"
func main() {
   s1 := `C:\sun\mon.tar.xz`
   // example 1
   s2 := filepath.Dir(s1)
   // example 2
   s3, s4 := filepath.Split(s1)
   // print
   println(s2 == `C:\sun`, s3 == `C:\sun\`, s4 == "mon.tar.xz")
}
