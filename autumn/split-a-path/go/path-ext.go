package main
import "path"

func main() {
   s := path.Ext("C:/go/README.md")
   println(s == ".md")
}
