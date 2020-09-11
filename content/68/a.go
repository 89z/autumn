package main
import "path"

func main() {
   s := "C:/go/bin/go.exe"
   s2 := path.Dir(s)
   println(s2 == "C:/go/bin")
}
