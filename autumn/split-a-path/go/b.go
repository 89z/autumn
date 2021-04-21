package main
import "path"

func main() {
   s := path.Dir("C:/go/README.md")
   println(s == "C:/go")
}
