package main
import "os"

func mkdir(s string) error {
   s, e := os.Stat(s)
   if e == nil && s.IsDir() { return nil }
   return os.Mkdir(s, os.ModeDir)
}

func main() {
   e := mkdir("north")
   if e != nil {
      panic(e)
   }
}
