package main
import "path"

func main() {
   file := "C:/go/bin/go.exe"
   dir := path.Dir(file)
   println(dir == "C:/go/bin")
}
