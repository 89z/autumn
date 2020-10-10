package main
import "os"

func main() {
   o, e := os.Stat("index.md")
   if e != nil {
      os.Exit(1)
   }
   b := o.Mode().IsRegular()
   println(b)
}
