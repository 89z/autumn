package main
import "os"

func main() {
   os.Chdir("..")
   s, e := os.Getwd()
   if e != nil {
      os.Exit(1)
   }
   println(s)
}
