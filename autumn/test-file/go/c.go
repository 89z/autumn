package main
import "os"

func main() {
   o, e := os.OpenFile("d.go", os.O_CREATE | os.O_EXCL, os.ModePerm)
   if os.IsExist(e) {
      os.Exit(1)
   }
   o.Close()
}
