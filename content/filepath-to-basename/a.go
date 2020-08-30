package main
import "path"
func main() {
   s1 := "C:/go/bin/go.exe"
   s2 := path.Base(s1)
   println(s2 == "go.exe")
}
