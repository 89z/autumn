package main
import "path"
func main() {
   s1 := "/sun/mon.tar.xz"
   // example 1
   s2 := path.Dir(s1)
   // example 2
   s3, s4 := path.Split(s1)
   // print
   println(s2 == "/sun", s3 == "/sun/", s4 == "mon.tar.xz")
}
