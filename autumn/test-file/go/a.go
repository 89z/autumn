package main
import "os"

func isFile(name string) bool {
   fi, e := os.Stat(name)
   return e == nil && fi.Mode().IsRegular()
}

func main() {
   b := isFile("a.go")
   println(b)
}
