package main
import "os"
func main() {
   o, e := os.Stat("index.md")
   // example 1
   if e != nil {
      os.Exit(1)
   }
   // example 2
   if os.IsNotExist(e) {
      os.Exit(1)
   }
   // example 3
   b1 := o.Mode().IsRegular()
   // example 4
   b2 := o.Mode().IsDir()
   // example 5
   b3 := o.IsDir()
   // print
   println(b1, b2, b3)
}
