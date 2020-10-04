package main
import "path"

func main() {
   s := "C:/go/bin/go.exe"
   s1 := path.Dir(s)
   println(s1 == "C:/go/bin")
}
