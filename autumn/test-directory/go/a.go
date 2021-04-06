package main
import "os"

func isDir(name string) bool {
   fi, err := os.Stat(name)
   return err == nil && fi.Mode().IsDir()
}

func main() {
   b := isDir(`C:\Users`)
   println(b)
}
