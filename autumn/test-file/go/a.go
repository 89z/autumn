package main
import "os"

func main() {
   o, e := os.Stat("index.md")
   if e != nil {
      os.Exit(1)
   }
   // example 1
   b1 := o.Mode().IsDir()
   // example 2
   b2 := o.IsDir()
   // print
   println(b1, b2)
}
