package main
import "path"
func main() {
   s1 := "/sunday/monday.tar.xz"
   // example 1
   s2 := path.Dir(s1)
   // example 2
   s3, _ := path.Split(s1)
   // print
   println(s2 == "/sunday", s3 == "/sunday/")
}
