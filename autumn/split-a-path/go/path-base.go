package main
import "path"

func main() {
   s := path.Base("C:/go/README.md")
   println(s == "README.md")
}
