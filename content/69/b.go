package main
import "os"
func main() {
   o, e := os.Stat("index.md")
   // example 1
   if os.IsNotExist(e) {
      os.Exit(1)
   }
   // example 2
   b1 := o.Mode().IsDir()
   // example 3
   b2 := o.Mode().IsRegular()
   // print
   println(b1, b2)
}
