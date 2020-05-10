package main
import "path"
func main() {
   s1 := "/sunday/monday.tar.xz"
   // example 1
   s2 := path.Base(s1)
   // example 2
   _, s3 := path.Split(s1)
   // example 3
   s4 := path.Ext(s1)
   // print
   println(s2 == "monday.tar.xz", s3 == "monday.tar.xz", s4 == ".xz")
}
