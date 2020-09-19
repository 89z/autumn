package main
import "path"

func main() {
   s := "C:/go/bin/go.exe"
   s1 := path.Base(s)
   println(s1 == "go.exe")
}
