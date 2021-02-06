package main
import "os"

func isFile(name string) bool {
   fi, err := os.Stat(name)
   return err == nil && fi.Mode().IsRegular()
}

func main() {
   b := isFile("a.go")
   println(b)
}
