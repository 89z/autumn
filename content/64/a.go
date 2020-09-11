package main
import "path"

func main() {
   s := "C:/go/bin/go.exe"
   s2 := path.Base(s)
   println(s2 == "go.exe")
}
