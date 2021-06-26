package main
import "os"

func isFile(name string) bool {
   info, err := os.Stat(name)
   return err == nil && ! info.Mode().IsDir()
}

func main() {
   b := isFile("file.txt")
   println(b)
}
