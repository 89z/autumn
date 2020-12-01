package main
import "os"

func main() {
   o, e := os.Stat("a.go")
   if e != nil {
      os.Exit(1)
   }
   n := o.Size()
   println(n)
}
