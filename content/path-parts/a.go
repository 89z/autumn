package main
import "path"
func main() {
   s1 := "/sunday/monday.tar.xz"
   // example 1
   b1 := path.Dir(s1) == "/sunday"
   // example 2
   b2 := path.Base(s1) == "monday.tar.xz"
   // example 3
   b3 := path.Ext(s1) == ".xz"
   // example 4
   s2, s3 := path.Split(s1)
   b4 := s2 == "/sunday/"
   b5 := s3 == "monday.tar.xz"
   // print
   println(b1, b2, b3, b4, b5)
}
