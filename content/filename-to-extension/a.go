package main
import "path"
func main() {
   s1 := "a.go"
   s2 := path.Ext(s1)
   println(s2 == ".go")
}
