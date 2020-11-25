package main
import "path"

func main() {
   in_s := "C:/go/bin/go.exe"
   out_s := path.Base(in_s)
   println(out_s == "go.exe")
}
