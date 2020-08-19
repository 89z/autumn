package main
import "path"
func main() {
   s := "/a/b.tar.xz"
   // example 1
   s1, s2 := path.Split(s)
   // example 2
   s3 := path.Base(s)
   // example 3
   s4 := path.Ext(s)
   // print
   println(s1 == "/a/", s2 == "b.tar.xz", s3 == "b.tar.xz", s4 == ".xz")
}
