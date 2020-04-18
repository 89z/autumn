package main
import "path"
func main() {
   s1 := "/sunday/monday.tar.xz"
   // example 1
   b1 := path.Base(s1) == "monday.tar.xz"
   // example 2
   b2 := path.Ext(s1) == ".xz"
   // example 3
   _, s2 := path.Split(s1)
   b3 := s2 == "monday.tar.xz"
   // print
   println(b1, b2, b3)
}
