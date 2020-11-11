package main
import "os"

func main() {
   o, e := os.Stat("a.go")
   if os.IsNotExist(e) || ! o.Mode().IsRegular() {
      os.Exit(1)
   }
}
