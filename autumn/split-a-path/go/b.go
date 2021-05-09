package main
import "path"

func main() {
   dir, base := path.Split("C:/go/README.md")
   println(dir == "C:/go/", base == "README.md")
}
