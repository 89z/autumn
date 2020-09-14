package main
import "os"

func main() {
   o, e := os.Stat("index.md")
   // example 1
   if os.IsNotExist(e) {
      os.Exit(1)
   }
   // example 2
   b := o.Mode().IsRegular()
   // example 3
   b2 := ! o.Mode().IsDir()
   // print
   println(b, b2)
}
