package main
import "path"
func main() {
   s1 := "C:/go/bin/go.exe"
   s2 := path.Dir(s1)
   println(s2 == "C:/go/bin")
}
