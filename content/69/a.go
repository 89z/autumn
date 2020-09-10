package main
import "os"
func main() {
   o, e := os.Stat("index.md")
   // example 1
   if e != nil {
      os.Exit(1)
   }
   // example 2
   b := o.IsDir()
   println(b)
}
