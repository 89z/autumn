package main
import "path"

func main() {
   s := path.Ext("a.go")
   println(s == ".go")
}
