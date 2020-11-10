package main
import "os"

func main() {
   o, e := os.Open("a.go")
   if e != nil {
      os.Exit(1)
   }
   o.Close()
}
