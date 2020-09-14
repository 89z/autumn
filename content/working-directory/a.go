package main
import "os"

func main() {
   s, e := os.Getwd()
   if e != nil {
      os.Exit(1)
   }
   println(s)
}
