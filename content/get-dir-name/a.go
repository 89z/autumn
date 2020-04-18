package main
import "path"
func main() {
   s1 := "/sunday/monday.tar.xz"
   // example 1
   b1 := path.Dir(s1) == "/sunday"
   // example 2
   s2, _ := path.Split(s1)
   b2 := s2 == "/sunday/"
   // print
   println(b1, b2)
}
