package main
import "os"

func main() {
   dir, e := os.ReadDir(".")
   if e != nil {
      panic(e)
   }
   for _, each := range dir {
      println(each.Name())
   }
}
