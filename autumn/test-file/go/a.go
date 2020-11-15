package main
import "os"

func IsFile(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.Mode().IsRegular()
}

func main() {
   b := IsFile("a.go")
   println(b)
}
