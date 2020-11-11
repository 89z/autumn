package main
import "os"

func main() {
   o, e := os.Open("a.go")
   if os.IsNotExist(e) {
      os.Exit(1)
   }
   o.Close()
}
